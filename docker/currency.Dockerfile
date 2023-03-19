# Builder
FROM golang:1.20.1-bullseye AS builder

WORKDIR /app/service

COPY ./currency .
COPY ./protos ../protos

RUN go mod download

RUN apt-get update && apt-get install -y --no-install-recommends make unzip

ENV PB_REL="https://github.com/protocolbuffers/protobuf/releases/download"
RUN curl -LO $PB_REL/v22.2/protoc-22.2-linux-x86_64.zip
RUN unzip protoc-22.2-linux-x86_64.zip -d /usr/local
ENV PATH="$PATH:/usr/local/bin"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED=0

RUN make protos
RUN make build

# Runner
FROM alpine:latest AS runner

COPY --from=builder /app/service/bin/server /app/server

CMD [ "/app/server" ]
