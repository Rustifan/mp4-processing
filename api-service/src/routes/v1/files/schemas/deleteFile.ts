import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";

const deleteFileParamsSchema = Type.Object({
    id: Type.Number()
})
const responseSchema = Type.Object({
    success: Type.Boolean(),
    message: Type.String()
})
const response = {
    200: responseSchema
}

export const deleteFileSchema: FastifySchema = {
    params: deleteFileParamsSchema,
    response
}

export type DeleteFileParams = Static<typeof deleteFileParamsSchema>
export type DeleteFileResponse = Static<typeof responseSchema>
export type DeleteFileTypes = {
    Params: DeleteFileParams,
    Resoponse: DeleteFileResponse
}