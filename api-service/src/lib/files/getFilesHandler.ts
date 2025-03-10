import { FastifyInstance } from "fastify";
import { FileDto } from "../../routes/v1/files/schemas/getFiles";

export type Props = Pick<FastifyInstance, "repositories">;

export const getFilesHandler = async ({ repositories }: Props): Promise<Array<FileDto>> => {
    return repositories.file.getAll();
};
