import { pgTable, text, timestamp, serial, pgEnum } from "drizzle-orm/pg-core";

export const fileStatusEnum = pgEnum('file_status', [
    'Processing',
    'Failed',
    'Successful',
]);

export const files = pgTable('files', {
    id: serial('id').primaryKey(),
    filePath: text('file_path').notNull(),
    status: fileStatusEnum('status').notNull(),
    processedFilePath: text('processed_file_path'),
    createdAt: timestamp('created_at').defaultNow(),
    updatedAt: timestamp('updated_at').defaultNow(),
});