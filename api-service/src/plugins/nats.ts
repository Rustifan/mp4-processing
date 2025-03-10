import { Type } from "@sinclair/typebox";
import { Static } from "@sinclair/typebox/";
import fp from "fastify-plugin";
import { connect, NatsConnection, PublishOptions, SubscriptionOptions } from "nats";
import { statusOptions } from "../config/constants";
import { getVerifiedTypeFromSchemaOrThrow } from "../lib/files/utils";

const NATS_TOPICS = {
    process_file: Type.Object({
        filePath: Type.String(),
    }),
    update_file: Type.Union([
        Type.Object({
            filePath: Type.String(),
            status: Type.Union(
                statusOptions
                    .filter((key) => key !== "Successful")
                    .map((option) => Type.Literal(option)),
            ),
        }),
        Type.Object({
            filePath: Type.String(),
            status: Type.Literal("Successful"),
            procssedFilePath: Type.String(),
        }),
    ]),
} as const;

export type NatsTopics = typeof NATS_TOPICS;
type NatsTopic = keyof typeof NATS_TOPICS;

class NatsClient {
    constructor(private nc: NatsConnection) {}
    public async publish<TTopic extends NatsTopic>(
        topic: TTopic,
        payload: Static<NatsTopics[TTopic]>,
        options?: PublishOptions,
    ): Promise<void> {
        const stringifiedPayload = JSON.stringify(payload);
        return this.nc.publish(topic, stringifiedPayload, options);
    }

    public async subscribe<TTopic extends NatsTopic>(
        topic: TTopic,
        onData: (data: Static<NatsTopics[TTopic]>) => void,
        onError?: (error: unknown) => void,
        options?: SubscriptionOptions,
    ) {
        const subOptions: SubscriptionOptions = {
            ...(options ?? {}),
            callback: (error, message) => {
                if (error && onError) {
                    return onError(error);
                }
                try {
                    const schema = NATS_TOPICS[topic];
                    const json = message.json();
                    const parsedData = getVerifiedTypeFromSchemaOrThrow(json, schema);
                    onData(parsedData);
                } catch (error) {
                    onError && onError(error);
                }
            },
        };
        this.nc.subscribe(topic, subOptions);
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
    { dependencies: ["environment"] },
);

declare module "fastify" {
    export interface FastifyInstance {
        natsClient: NatsClient;
    }
}
