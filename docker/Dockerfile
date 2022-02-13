FROM golang:1.17-alpine AS build-env

WORKDIR /build

COPY . .

RUN go build -o /build/cmd/app /build/cmd/main.go

#
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=build-env /build/cmd/app /app/main

WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main","app"]
