import { FastifyPluginAsync } from 'fastify'

const root: FastifyPluginAsync = async (fastify, opts): Promise<void> => {
  fastify.get('/', async function (request, reply) {
    const results = await fastify.repositories.file.getAll();
    console.log(results)

    return { root: true, files: results }
  })
}

export default root;
