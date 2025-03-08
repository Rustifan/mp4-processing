import { BasicRepoProps } from "."
import { files } from "../schema/files"

export type File = typeof files.$inferSelect

export class FileRepository {
    constructor(private dependencies: BasicRepoProps) { }
    async getAll(): Promise<Array<File>> {
        return await this.dependencies.db.select().from(files).execute()
    }
}



