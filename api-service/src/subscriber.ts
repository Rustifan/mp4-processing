import fastify, { FastifyInstance } from "fastify";
import { join } from "path";
import AutoLoad from "@fastify/autoload";
import { updateFileStatus } from "./lib/files/updateFileStatusHandler";

function readyPromise(app: FastifyInstance): Promise<void> {
    return new Promise((resolve, reject) => {
        app.ready((error) => {
            if (error) {
                reject(error);
            } else {
                resolve(undefined);
            }
        });
    });
}

async function main() {
    const app: FastifyInstance = fastify({ logger: true });
    app.register(AutoLoad, {
        dir: join(__dirname, "plugins"),
        options: {},
    });

    await readyPromise(app);
    app.log.info("Plugins registered");
    app.natsClient.subscribe(
        "update_file",
        (data) => updateFileStatus(app, data),
        (error) => {
            app.log.error("Error happened while parsing subscription message")
            console.log(error)
        },
        { queue: "update_file_queue" }
    );
}

main();
