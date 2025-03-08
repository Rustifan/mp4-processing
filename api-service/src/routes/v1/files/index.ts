import { FastifyPluginAsync } from 'fastify'
import { getFilesSchema } from './schemas/getFiles';
import { getFilesHandler } from '../../../lib/files/getFilesHandler';
import { startProcessingSchema, StartProcessingTypes } from './schemas/postStartProcessing';
import { startProcessingHandler } from '../../../lib/files/startProcessingHandler';
import { deleteFileSchema, DeleteFileTypes } from './schemas/deleteFile';
import { deleteFileHandler } from '../../../lib/files/deleteFileHandler';


const file: FastifyPluginAsync = async (fastify): Promise<void> => {
    fastify.get('/', { schema: getFilesSchema }, () => getFilesHandler(fastify))
    fastify.post<StartProcessingTypes>("/start-processing", { schema: startProcessingSchema }, (request) => startProcessingHandler(fastify, request.body.filePath))
    fastify.delete<DeleteFileTypes>("/:id", { schema: deleteFileSchema }, (request) => deleteFileHandler(fastify, request.params.id))
}

export default file;
