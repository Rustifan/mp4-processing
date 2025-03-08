import { join } from "path"
import { FILE_FOLDER, PROCESSED_FILES_FOLDER } from "../../config/constants"

export function getFilePathInFilesFolder(filePath: string) {
    return join(FILE_FOLDER, filePath)
}

export function getFilePathInProcessedFilesFolder(filePath: string) {
    return join(PROCESSED_FILES_FOLDER, filePath)
}