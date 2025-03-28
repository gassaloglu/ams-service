FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags '-w -s -extldflags "-static"' -o build/main -a cmd/main.go

FROM scratch
COPY --from=builder /app/build ./build
ENTRYPOINT [ "./build/main" ]