/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    API_HOSTNAME: process.env.API_HOSTNAME || 'api', // Assurez-vous que l'hôte API est correct
  },
}

module.exports = nextConfig
