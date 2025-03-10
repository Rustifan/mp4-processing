import { FastifyInstance } from "fastify";
import { files } from "../schema/files";
import { FileRepository } from "./file";

export type BasicRepoProps = Pick<FastifyInstance, "db">;
export type File = typeof files.$inferSelect;

export type Repositories = {
    file: FileRepository;
};

export function getRepositories(dependencies: BasicRepoProps) {
    return {
        file: new FileRepository(dependencies),
    };
}
