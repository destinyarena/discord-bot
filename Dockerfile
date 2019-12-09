FROM golang:1.13-buster AS builder

WORKDIR /build
COPY . .
RUN apt update && apt install -y build-essential
RUN make build

FROM scratch
WORKDIR /app
COPY --from=builder /build/bin/* /app

CMD ["./d2arena"]
