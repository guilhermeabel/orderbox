FROM golang:alpine AS build

RUN apk update && apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -ldflags="-w -s" -o /bin/main /app/cmd/web

FROM alpine

RUN apk update && apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/main /


ENTRYPOINT ["/main"]
