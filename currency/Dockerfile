FROM rust:1.68-slim AS builder

ARG COMMIT_SHA=""
ARG VERSION=""

WORKDIR /app

# Install and build dependencies to prevent cache invalidation
# COPY ./Cargo.toml ./Cargo.lock ./
# RUN mkdir src && echo "fn main() {}" > src/main.rs 
# RUN cargo build --release && rm src/*.rs

# Insert commit sha and version
RUN mkdir .meta
RUN echo ${COMMIT_SHA} > .meta/HASH
RUN echo ${VERSION} > .meta/VERSION

# Copy source code and build
COPY . .
RUN cargo build --release

# Runtime image
FROM debian:buster-slim AS runtime

WORKDIR /app

COPY --from=builder /app/.meta ./.meta
COPY --from=builder /app/log4rs.yaml ./log4rs.yaml
COPY --from=builder /app/target/release/service-currency ./service-currency

CMD ["./service-currency"]