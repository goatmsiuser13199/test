# Stage 1: Build the Next.js application
FROM node:20-alpine AS builder
WORKDIR /app
COPY package-lock.json .
COPY package.json .
RUN npm ci
COPY . .
RUN npx next build

# Stage 2: Create a lightweight image
FROM node:20-alpine
ENV API_HOSTNAME=api
WORKDIR /app
COPY --from=builder /app ./
CMD [ "npx", "next", "start", "-p", "80" ]