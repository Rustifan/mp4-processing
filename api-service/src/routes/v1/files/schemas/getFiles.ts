import { Static, Type } from "@sinclair/typebox";
import { FastifySchema } from "fastify";
import { FileStatus, statusOptions } from "../../../../config/constants";

const statusEnum = statusOptions.reduce((acc, curr) => ({ ...acc, [curr]: curr }), {} as { [K in FileStatus]: K })
const fileDtoSchema = Type.Object({
    id: Type.Number({ description: "Id of file record" }),
    filePath: Type.String({ description: "File path" }),
    status: Type.Enum(statusEnum, { description: "Status of processing" }),
    processedFilePath: Type.Union([Type.String({ description: "Path of processed file" }), Type.Null()])
})

const filesDtoSchema = Type.Array(fileDtoSchema)

export type FileDto = Static<typeof fileDtoSchema>

export const getFilesSchema: FastifySchema = {
    response: {
        200: filesDtoSchema
    }
}