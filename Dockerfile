# Create Builder Image, to compile the source code into an executable
FROM golang:1.22.0-alpine as builder
RUN pwd 
RUN ls -la
RUN apk add --no-cache git gcc musl-dev
RUN go mod download
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Create the final image, running the API and exposing port 8080
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
ARG PORT
ENV PORT=$PORT
EXPOSE $PORT
CMD ["./main"]