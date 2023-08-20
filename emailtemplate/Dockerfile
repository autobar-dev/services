# Builder image
FROM node:20 AS builder

ARG GITHUB_USER
ARG GITHUB_PAT

RUN apt update && apt install git
RUN echo "machine github.com login $GITHUB_USER password $GITHUB_PAT" > ~/.netrc

WORKDIR /app

# Insert commit sha and version
ARG COMMIT_SHA=""
ARG VERSION=""

RUN mkdir .meta
RUN echo ${COMMIT_SHA} > .meta/HASH
RUN echo ${VERSION} > .meta/VERSION

# Copy package.json and package-lock and install dependencies
COPY package*.json ./
RUN npm ci

# Copy source code and build
COPY . .
RUN npm run build

# Runtime image
FROM node:20-alpine AS runtime

WORKDIR /app

COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/.meta ./.meta
COPY --from=builder /app/dist ./dist

CMD ["node", "dist/main.js"]
