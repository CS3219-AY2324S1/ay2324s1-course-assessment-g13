/** @type {import('next').NextConfig} */
// eslint-disable-next-line no-undef
module.exports = {
  output: 'standalone',
  webpack: config => {
    config.watchOptions = {
      poll: 1000,
      aggregateTimeout: 300,
    };
    return config;
  },
  images: {
    domains: ['res.cloudinary.com'],
  },
};
