# Stage 1
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./server

# Stage 2
FROM scratch

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 50051

ENTRYPOINT ["/app/server"]