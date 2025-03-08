import { QueryResult } from "pg"
import { BasicRepoProps } from "."
import { files } from "../schema/files"
import { eq } from "drizzle-orm"

export type File = typeof files.$inferSelect

export class FileRepository {
    constructor(private dependencies: BasicRepoProps) { }
    async getAll(): Promise<Array<File>> {
        return await this.dependencies.db.select().from(files).execute()
    }

    async insertFile(filePath: string, status: File["status"] = "Processing"): Promise<QueryResult> {
        return this.dependencies.db.insert(files).values({ filePath, status })
    }

    async deleteFileHandler(id: number): Promise<QueryResult<never>> {
        return this.dependencies.db.delete(files).where(eq(files.id, id))
    }

}



