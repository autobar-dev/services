# Dependency builder image
FROM rust:1.69-slim AS deps-builder

WORKDIR /app

RUN apt-get update && apt-get install -y pkg-config libssl-dev

# Install and build dependencies to prevent cache invalidation
COPY ./Cargo.toml ./Cargo.lock ./
RUN mkdir src && echo "fn main() { println!(\"dummy\"); }" > src/main.rs 
RUN cargo build --release && rm -rf src/ && rm -f target/release/service-currency*

# App dependency builder
FROM deps-builder AS builder

# Insert commit sha and version
ARG COMMIT_SHA=""
ARG VERSION=""

RUN mkdir .meta
RUN echo ${COMMIT_SHA} > .meta/HASH
RUN echo ${VERSION} > .meta/VERSION

# Copy source code and build
COPY . .
RUN cargo clean --release -p service-currency && cargo build --release

# Runtime image
FROM debian:buster-slim AS runtime

RUN apt update && apt install -y libssl-dev ca-certificates

WORKDIR /app

COPY --from=builder /app/.meta ./.meta
COPY --from=builder /app/log4rs.yaml ./log4rs.yaml
COPY --from=builder /app/target/release/service-currency ./service-currency

CMD ["./service-currency"]
