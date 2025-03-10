import { FastifyInstance } from "fastify";
import { NatsTopics } from "../../plugins/nats";
import { Static } from "@sinclair/typebox";

type Props = Pick<FastifyInstance, "repositories" | "log">;

export async function updateFileStatus(
    { repositories, log }: Props,
    payload: Static<NatsTopics["update_file"]>,
) {
    const processedFilePath =
        payload.status === "Successful" ? payload.procssedFilePath : undefined;
    const message = payload.status !== "Successful" ? payload.message : undefined;
    try {
        const updateResult = await repositories.file.updateFileStatus(
            payload.filePath,
            payload.status,
            processedFilePath,
            message,
        );
        updateResult.rowCount
            ? log.info(`Updated file with path ${payload.filePath}`)
            : log.error(`File with path ${payload.filePath} not updated`);
    } catch (error) {
        log.error("Something went wrong");
    }
}
