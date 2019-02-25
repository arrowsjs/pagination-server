## Builder
FROM golang:alpine3.8 AS builder
RUN apk add --no-cache git
WORKDIR /build
COPY ./ /build
RUN CGO_ENABLED=0 go build

## Application
FROM scratch
CMD ["./pagination-server"]
COPY data ./data
COPY --from=builder /build/pagination-server ./pagination-server
