import fp from "fastify-plugin";
import { Type, Static } from "@sinclair/typebox";
import { envSchema } from "env-schema";

const schema = Type.Object({
    DATABASE_URL: Type.String(),
    NATS_URL: Type.String(),
    PORT: Type.String(),
});
type Environment = Static<typeof schema>;

export default fp(
    async (fastify) => {
        const environement = envSchema<Environment>({
            schema,
            data: process.env,
            dotenv: true,
        });

        fastify.decorate("environment", environement);
        fastify.log.info("Environment loaded");
    },
    { name: "environment" },
);

declare module "fastify" {
    export interface FastifyInstance {
        environment: Environment;
    }
}
