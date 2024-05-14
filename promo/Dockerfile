# Builder image
FROM golang:alpine AS builder

ARG GITHUB_USER
ARG GITHUB_PAT

RUN apk update && apk add git
RUN echo "machine github.com login $GITHUB_USER password $GITHUB_PAT" > ~/.netrc

WORKDIR /app

# Insert commit sha and version
ARG COMMIT_SHA=""
ARG VERSION=""

RUN mkdir .meta
RUN echo ${COMMIT_SHA} > .meta/HASH
RUN echo ${VERSION} > .meta/VERSION

# Copy source code and build
COPY . .
ENV GOPRIVATE=github.com/autobar-dev/shared-libraries
RUN go build -o ./service-promo

# Runtime image
FROM alpine AS runtime

WORKDIR /app

COPY --from=builder /app/.meta ./.meta
# COPY --from=builder /app/log4rs.yaml ./log4rs.yaml
COPY --from=builder /app/service-promo ./service-promo

CMD ["./service-promo"]
