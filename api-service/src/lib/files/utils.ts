import { join } from "path";
import { FILE_FOLDER, PROCESSED_FILES_FOLDER } from "../../config/constants";
import { Static, TObject, TUnion } from "@sinclair/typebox";
import { TypeCompiler } from "@sinclair/typebox/compiler";

export function getFilePathInFilesFolder(filePath: string) {
    return join(FILE_FOLDER, filePath);
}

export function getFilePathInProcessedFilesFolder(filePath: string) {
    return join(PROCESSED_FILES_FOLDER, filePath);
}

export function getVerifiedTypeFromSchemaOrThrow<TSchema extends TObject | TUnion>(
    object: unknown,
    schema: TSchema,
): Static<TSchema> {
    const compiledSchema = TypeCompiler.Compile(schema);
    if (!compiledSchema.Check(object)) {
        const message = `Object does not match schema. Errors: ${JSON.stringify(Array.from(compiledSchema.Errors(object)), null, 2)}`;
        throw new Error(message);
    }

    return object as Static<TSchema>;
}
