#build stage
FROM golang:alpine AS builder
ENV AWS_ACCESS_KEY_ID=AKIATAXQ52ELAXRNFGH5
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_REGION=us-east-1
ENV GIN_MODE=release
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/app -v .

#final stage
FROM alpine:latest
RUN apk update && apk upgrade
RUN apk add --no-cache sqlite

ENV AWS_ACCESS_KEY_ID=AKIATAXQ52ELAXRNFGH5
ENV AWS_SECRET_ACCESS_KEY=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
ENV AWS_REGION=us-east-1
ENV GIN_MODE=release
COPY --from=builder /go/bin/app /usr/bin/
RUN mkdir /app
WORKDIR /app
EXPOSE 8080

ENTRYPOINT ["/usr/bin/app"] 
