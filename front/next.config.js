/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    API_HOSTNAME: process.env.API_HOSTNAME || 'localhost',
  },
}

module.exports = nextConfig
