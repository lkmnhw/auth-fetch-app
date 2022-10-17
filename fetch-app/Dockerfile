FROM golang:1.19.2-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/fetch-app
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/fetch-app

FROM golang:1.19.2-alpine
COPY --from=builder /go/bin/fetch-app /go/bin/fetch-app

RUN addgroup go && adduser -D -G go go \
    && chown -R go:go /go/bin/fetch-app
USER go

EXPOSE 4000
