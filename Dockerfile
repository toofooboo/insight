FROM golang:1.23-alpine AS build

RUN apk add --no-cache git bash

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/api cmd/api/main.go
RUN go build -o /bin/indexer cmd/indexer/main.go

FROM alpine:3.18

RUN apk add --no-cache bash

COPY --from=build /bin/api /bin/api
COPY --from=build /bin/indexer /bin/indexer

EXPOSE 8080

CMD ["sh", "-c", "/bin/api & /bin/indexer"]