import { Type } from "@sinclair/typebox";
import { Static } from "@sinclair/typebox/";
import fp from "fastify-plugin";
import { connect, NatsConnection, PublishOptions } from "nats";

const NATS_TOPICS = {
    process_file: Type.Object({
        filePath: Type.String(),
    }),
} as const;

type NatsTopics = typeof NATS_TOPICS;
type NatsTopic = keyof typeof NATS_TOPICS;

class NatsClient {
    constructor(private nc: NatsConnection) { }
    public async publish<TTopic extends NatsTopic>(
        topic: TTopic,
        payload: Static<NatsTopics[TTopic]>,
        options?: PublishOptions
    ): Promise<void> {
        const stringifiedPayload = JSON.stringify(payload);
        return this.nc.publish(topic, stringifiedPayload, options);
    }
}

export default fp(
    async (fastify) => {
        const nc = await connect({
            servers: fastify.environment.NATS_URL,
        });
        const natsClient = new NatsClient(nc);
        fastify.decorate("natsClient", natsClient);
    },
    { dependencies: ["environment"] }
);

declare module "fastify" {
    export interface FastifyInstance {
        natsClient: NatsClient;
    }
}
