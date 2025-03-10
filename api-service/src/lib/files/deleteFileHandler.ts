import { FastifyInstance } from "fastify";
import { DeleteFileResponse } from "../../routes/v1/files/schemas/deleteFile";
import { getFilePathInProcessedFilesFolder } from "./utils";

type Props = Pick<FastifyInstance, "httpErrors" | "repositories" | "diskOperations" | "log">;

export async function deleteFileHandler(
    { httpErrors, repositories, diskOperations, log }: Props,
    fileId: number,
): Promise<DeleteFileResponse> {
    const file = await repositories.file.getById(fileId);
    if (!file) {
        throw httpErrors.notFound(`File with id ${fileId} not found`);
    }
    const result = await repositories.file.deleteFileHandler(fileId);
    if (!result.rowCount || result.rowCount < 1) {
        throw httpErrors.internalServerError("Something went wrong while deleting file");
    }
    log.info(`File with id ${fileId} deleted`);

    if (file.processedFilePath) {
        await diskOperations
            .deleteFile(getFilePathInProcessedFilesFolder(file.processedFilePath))
            .then(() => log.info(`Deleted processed file on path ${file.processedFilePath}`))
            .catch((error) => {
                log.error("Error while deleteing file");
                log.error(error);
            });
    }

    return {
        message: "File information deleted",
        success: true,
    };
}
