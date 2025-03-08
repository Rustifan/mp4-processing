import { join } from "path"
import { FILE_FOLDER } from "../../config/constants"

export function getFilePathInFilesFolder(filePath: string) {
    return join(FILE_FOLDER, filePath)
}