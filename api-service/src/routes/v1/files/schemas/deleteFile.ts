import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";
import { errorResponseSchema } from ".";

const deleteFileParamsSchema = Type.Object({
    id: Type.Number({
        description: "Id of a file process that you want to delete",
    }),
});
const responseSchema = Type.Object({
    success: Type.Boolean({ description: "Success flag" }),
    message: Type.String({ description: "Message description" }),
});
const response = {
    200: responseSchema,
    default: errorResponseSchema,
};

export const deleteFileSchema: FastifySchema = {
    params: deleteFileParamsSchema,
    response,
    tags: ["delete-file"],
};

export type DeleteFileParams = Static<typeof deleteFileParamsSchema>;
export type DeleteFileResponse = Static<typeof responseSchema>;
export type DeleteFileTypes = {
    Params: DeleteFileParams;
    Resoponse: DeleteFileResponse;
};
