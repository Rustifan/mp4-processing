import fp from 'fastify-plugin'
import { drizzle } from 'drizzle-orm/node-postgres';
import { getRepositories, Repositories } from '../db/repository';

export default fp(async (fastify) => {
    const db = drizzle(fastify.environment.DATABASE_URL)
    fastify.decorate('db', db)
    fastify.decorate('repositories', getRepositories(fastify))

}, { dependencies: ["environment"] })


declare module "fastify" {
    export interface FastifyInstance {
        db: ReturnType<typeof drizzle>,
        repositories: Repositories
    }
}