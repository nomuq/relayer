FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o relayer cmd/server/main.go

FROM golang:alpine
COPY --from=builder /app/relayer /relayer
COPY --from=builder /app/sql sql
RUN mkdir -p /etc/relayer
EXPOSE 1203
CMD ["/relayer", "--config", "/etc/relayer/config.yaml"]
