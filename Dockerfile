FROM golang:alpine3.12 AS build

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/test ./main.go

FROM alpine:3.10
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8090
EXPOSE 8091
ENTRYPOINT /go/bin/test