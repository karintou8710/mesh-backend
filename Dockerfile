FROM golang:1.22.5 AS go

WORKDIR /app
COPY ./ ./
RUN go mod download && go build -o main server/*.go

EXPOSE 8080

CMD ["/app/main"]