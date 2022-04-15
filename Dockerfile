FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o relayer cmd/relayer/main.go

FROM golang:alpine
COPY --from=builder /app/relayer /relayer
RUN mkdir -p /etc/relayer
EXPOSE 1203
CMD ["/relayer", "--config", "/etc/relayer/config.yaml"]
