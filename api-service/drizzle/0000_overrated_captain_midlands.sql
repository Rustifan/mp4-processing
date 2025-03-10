CREATE TYPE "public"."file_status" AS ENUM('Processing', 'Failed', 'Successful');--> statement-breakpoint
CREATE TABLE "files" (
	"id" serial PRIMARY KEY NOT NULL,
	"file_path" text NOT NULL,
	"status" "file_status" NOT NULL,
	"processed_file_path" text,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp DEFAULT now()
);
--> statement-breakpoint
CREATE UNIQUE INDEX "file_path_idx" ON "files" USING btree ("file_path");