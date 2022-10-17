FROM golang:1.19.2-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/auth-app
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/auth-app

FROM golang:1.19.2-alpine
COPY --from=builder /go/bin/auth-app /go/bin/auth-app

RUN addgroup go && adduser -D -G go go \
    && chown -R go:go /go/bin/auth-app
USER go

EXPOSE 3000
