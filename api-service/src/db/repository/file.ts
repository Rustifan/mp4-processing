import { QueryResult } from "pg";
import { BasicRepoProps } from ".";
import { files } from "../schema/files";
import { eq } from "drizzle-orm";
import { FileStatus } from "../../config/constants";

export type File = typeof files.$inferSelect;

export class FileRepository {
    constructor(private dependencies: BasicRepoProps) {}
    async getAll(): Promise<Array<File>> {
        return this.dependencies.db.select().from(files).execute();
    }

    async getById(id: number): Promise<File | null> {
        const results = await this.dependencies.db.select().from(files).where(eq(files.id, id));
        return results.at(0) ?? null;
    }

    async insertFile(
        filePath: string,
        status: File["status"] = "Processing",
    ): Promise<QueryResult> {
        return this.dependencies.db.insert(files).values({ filePath, status });
    }

    async deleteFileHandler(id: number): Promise<QueryResult<never>> {
        return this.dependencies.db.delete(files).where(eq(files.id, id));
    }

    async updateFileStatus(
        filePath: string,
        status: FileStatus,
        processedFilePath?: string,
        message?: string,
    ): Promise<QueryResult<never>> {
        return this.dependencies.db
            .update(files)
            .set({ status, processedFilePath, message })
            .where(eq(files.filePath, filePath));
    }
}
