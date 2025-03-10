import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";
import { errorResponseSchema } from ".";

const bodySchema = Type.Object({
    filePath: Type.String({
        description:
            "File path relative to file folder. If there are not nested folders you can use only file name",
    }),
});

const responseSchema = Type.Object({
    success: Type.Boolean({ description: "Success flag" }),
    message: Type.String({ description: "Message description" }),
});
export const response = {
    200: responseSchema,
    default: errorResponseSchema,
};
export type StartProcessingBody = Static<typeof bodySchema>;
export const startProcessingSchema: FastifySchema = {
    body: bodySchema,
    response,
    tags: ["start-processing"],
    description: "Start processing file",
};
export type StartProcessingResponse = Static<typeof responseSchema>;
export type StartProcessingTypes = {
    Body: StartProcessingBody;
    Reseponse: StartProcessingResponse;
};
