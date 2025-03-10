import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";
import { FileStatus, statusOptions } from "../../../../config/constants";
import { errorResponseSchema } from ".";

const statusEnum = statusOptions.reduce(
    (acc, curr) => ({ ...acc, [curr]: curr }),
    {} as { [K in FileStatus]: K },
);
const fileDtoSchema = Type.Object({
    id: Type.Number({ description: "Id of file record" }),
    filePath: Type.String({ description: "File path" }),
    status: Type.Enum(statusEnum, { description: "Status of processing" }),
    processedFilePath: Type.Union([Type.String(), Type.Null()], {
        description: "Path of processed file",
    }),
    message: Type.Union([Type.String(), Type.Null()], {
        description: "Status message",
    }),
});

const filesDtoSchema = Type.Array(fileDtoSchema);

export type FileDto = Static<typeof fileDtoSchema>;

export const getFilesSchema: FastifySchema = {
    response: {
        200: filesDtoSchema,
        default: errorResponseSchema,
    },
    tags: ["get-all-files"],
    description: "Return all file processings",
};
