FROM rust:1.68-slim AS builder

WORKDIR /app

# Install and build dependencies to prevent cache invalidation
# COPY ./Cargo.toml ./Cargo.lock ./
# RUN mkdir src && echo "fn main() {}" > src/main.rs 
# RUN cargo build --release && rm src/*.rs

# Copy source code and build
COPY . .
RUN cargo build --release

# Runtime image
FROM debian:buster-slim AS runtime

WORKDIR /app

COPY --from=builder /app/log4rs.yaml ./log4rs.yaml
COPY --from=builder /app/target/release/service-currency ./service-currency

CMD ["./service-currency"]