FROM golang:1.18-alpine AS builder

WORKDIR /family

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o family main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /family

COPY --from=builder /family/family ./
COPY ./docker/local/family.sh ./
COPY ./core-contracts/artifacts ./core-contracts/artifacts

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./family.sh"]
