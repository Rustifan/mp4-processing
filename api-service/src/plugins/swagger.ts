import fp from "fastify-plugin";
import { FastifySensibleOptions } from "@fastify/sensible";
import swagger from "@fastify/swagger";
import swaggerUi from "@fastify/swagger-ui";

export default fp<FastifySensibleOptions>(async (fastify) => {
    await fastify.register(swagger, {
        swagger: {
            info: {
                title: "Mp4-api",
                description: "Mp4 processing api",
                version: "0.1.0",
            },
        },
        hideUntagged: true,
    });
    await fastify.register(swaggerUi, {
        routePrefix: "/swagger",
    });
});
