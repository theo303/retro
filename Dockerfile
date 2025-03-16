FROM node:23-alpine AS node-builder

WORKDIR /front

COPY front .

RUN npm install
RUN npm run build

FROM golang:1.23 AS go-builder

WORKDIR /go/src/github.com/theo303/retro

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd/server cmd/server/

RUN CGO_ENABLED=0 GOOS=linux go build -o server.out ./cmd/server

FROM alpine:latest

WORKDIR /

COPY --from=go-builder /go/src/github.com/theo303/retro/server.out .
COPY --from=node-builder /front/dist public

EXPOSE 8080

CMD ["./server.out"]
