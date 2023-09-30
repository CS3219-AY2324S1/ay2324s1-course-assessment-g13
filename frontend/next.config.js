/** @type {import('next').NextConfig} */
// eslint-disable-next-line no-undef
module.exports = {
  output: 'standalone',
  webpack: (config, _) => ({
    ...config,
    watchOptions: {
      ...config.watchOptions,
      poll: 1000,
      aggregateTimeout: 300,
    },
  }),
};
