import { FastifyInstance } from "fastify";
import { NatsTopics } from "../../plugins/nats";
import { Static } from "@sinclair/typebox";

type Props = Pick<FastifyInstance, "repositories" | "log">;

export async function updateFileStatus(
    { repositories, log }: Props,
    payload: Static<NatsTopics["update_file"]>
) {
    const processedFilePath =
        payload.status === "Successful" ? payload.procssedFilePath : undefined;
    try {
        const updateResult = await repositories.file.updateFileStatus(
            payload.id,
            payload.status,
            processedFilePath
        );
        updateResult.rowCount
            ? log.info(`Updated file with id ${payload.id}`)
            : log.error(`File with id ${payload.id} not updated`);
    } catch (error) {
        log.error("Something went wrong");
    }
}
