import fp from "fastify-plugin";
import fs from "fs/promises";

async function fileExists(filePath: string) {
    try {
        await fs.access(filePath, fs.constants.F_OK);
        return true;
    } catch (err) {
        return false;
    }
}

async function deleteFile(filePath: string) {
    return fs.rm(filePath);
}

const diskOperations = {
    fileExists,
    deleteFile,
} as const;

export default fp(async (fastify) => {
    fastify.decorate("diskOperations", diskOperations);
});

declare module "fastify" {
    export interface FastifyInstance {
        diskOperations: typeof diskOperations;
    }
}
