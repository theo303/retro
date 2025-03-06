FROM golang:1.23 AS builder

ARG BUILD_DATE
ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/github.com/theo303/retro

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd/server cmd/server/

RUN CGO_ENABLED=0 GOOS=linux go build -o server.out ./cmd/server

FROM alpine:latest

WORKDIR /

COPY --from=builder /go/src/github.com/theo303/retro/server.out .
COPY public public

EXPOSE 8080

CMD ["./server.out"]
