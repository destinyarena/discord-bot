FROM golang:alipne AS builder

WORKDIR /build
COPY . .
RUN apk add --update make
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/bin/d2arena /app

CMD ["./d2arena"]
