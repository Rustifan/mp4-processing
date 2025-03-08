import fp from 'fastify-plugin'
import fs from "fs/promises"

async function fileExists(filePath: string) {
    try {
        await fs.access(filePath, fs.constants.F_OK);
        return true;
    } catch (err) {
        return false;
    }
}

const diskOperations = {
    fileExists
} as const

export default fp(async (fastify) => {
    fastify.decorate("diskOperations", diskOperations)
})


declare module 'fastify' {
    export interface FastifyInstance {
        diskOperations: typeof diskOperations
    }
}
