import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";

const bodySchema = Type.Object({
    filePath: Type.String()
})

const responseSchema = Type.Object({
    success: Type.Boolean(),
    message: Type.String()
})
export const response = {
    200: responseSchema
}
export type StartProcessingBody = Static<typeof bodySchema>
export const startProcessingSchema: FastifySchema = {
    body: bodySchema,
    response
}
export type StartProcessingResponse = Static<typeof responseSchema>
export type StartProcessingTypes = {
    Body: StartProcessingBody,
    Reseponse: StartProcessingResponse
}
