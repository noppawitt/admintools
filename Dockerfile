FROM golang:1.11.1-alpine3.8 AS builder

RUN apk add git

WORKDIR /go/src/github.com/noppawitt/admintools

ENV GO111MODULE on

# ENV APP_ENV development

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s' -o /go/bin/admintools

FROM scratch

WORKDIR /go/bin

COPY --from=builder /go/bin/admintools /go/src/github.com/noppawitt/admintools/config/*.json ./

ENTRYPOINT ["./admintools"]
