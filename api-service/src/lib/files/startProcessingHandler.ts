import { FastifyInstance } from "fastify"
import { StartProcessingResponse } from "../../routes/v1/files/schemas/postStartProcessing"
import { getFilePathInFilesFolder } from "./utils"

export type Props = Pick<FastifyInstance, "repositories" | "httpErrors" | "diskOperations">

export async function startProcessingHandler({ repositories, httpErrors, diskOperations }: Props, filePath: string): Promise<StartProcessingResponse> {
    const pathInFilesFolder = getFilePathInFilesFolder(filePath);
    const fileExistsInFolder = await diskOperations.fileExists(pathInFilesFolder)
    if (!fileExistsInFolder) {
        throw httpErrors.badRequest(`File does not exist in path ${filePath}`)
    }
    const result = await repositories.file.insertFile(filePath)
    if (!result.rowCount || result.rowCount < 1) {
        throw httpErrors.internalServerError("Something went wrong while saving data to db")
    }

    return {
        success: true,
        message: "Processing started"
    }
}