import { FastifyInstance } from "fastify";
import { DeleteFileResponse } from "../../routes/v1/files/schemas/deleteFile";

type Props = Pick<FastifyInstance, "httpErrors" | "repositories">

export async function deleteFileHandler({ httpErrors, repositories }: Props, fileId: number): Promise<DeleteFileResponse> {
    const result = await repositories.file.deleteFileHandler(fileId)
    if (!result.rowCount || result.rowCount < 1) {
        throw httpErrors.notFound(`File with id ${fileId} not found`)
    }

    return {
        message: "File information deleted",
        success: true
    }
}