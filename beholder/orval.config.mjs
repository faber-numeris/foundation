export default {
  api: {
    input: './openapi.yaml',
    output: {
      mode: 'tags-split',
      target: './ui/src/api',
      schemas: './ui/src/models',
      client: 'fetch',
    },
  },
  apiZod:{
    input: './openapi.yaml',
    output: {
      mode: 'tags-split',
      target: './ui/src/api',
      schemas: './ui/src/models',
      client: 'zod',
      override: {
        zod: {
          version: 4,
        },
      },
    },
  }
};
