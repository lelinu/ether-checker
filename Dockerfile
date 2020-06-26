FROM golang:alpine as builder
RUN mkdir /build
COPY ./src /build/src
COPY go.mod go.sum /build/

WORKDIR /build/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/src/main /app/
WORKDIR /app
CMD ["./main"]
