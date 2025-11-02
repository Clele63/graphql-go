ARG NODE_VERSION=20
ARG GO_VERSION=1.25
ARG ALPINE_BUILD_VERSION=3.20

FROM node:${NODE_VERSION}-alpine AS client-builder

WORKDIR /app/client

COPY client/package*.json ./
RUN npm install

COPY client/ ./
RUN npm run build


FROM golang:${GO_VERSION}-alpine AS server-builder

WORKDIR /app/server

RUN apk add --no-cache git

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server server.go


FROM alpine:${ALPINE_BUILD_VERSION} AS graphql-go-app

WORKDIR /app

COPY --from=server-builder /server /server
COPY --from=client-builder /app/client/build ./client

RUN mkdir -p /data && chmod 777 /data

CMD ["/server"]
