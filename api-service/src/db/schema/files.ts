import { pgTable, text, timestamp, serial, pgEnum } from "drizzle-orm/pg-core";
import { statusOptions } from "../../config/constants";
import { uniqueIndex } from "drizzle-orm/pg-core";

export const fileStatusEnum = pgEnum("file_status", statusOptions);

export const files = pgTable(
    "files",
    {
        id: serial("id").primaryKey(),
        filePath: text("file_path").notNull(),
        status: fileStatusEnum("status").notNull(),
        processedFilePath: text("processed_file_path"),
        createdAt: timestamp("created_at").defaultNow(),
        updatedAt: timestamp("updated_at").defaultNow(),
        message: text("message"),
    },
    (table) => [uniqueIndex("file_path_idx").on(table.filePath)],
);
