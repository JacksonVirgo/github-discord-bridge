FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ARG ENV_FILE
ENV ENV_FILE=${ENV_FILE}
COPY $ENV_FILE .env
COPY . .
RUN go build -o ./bin/main src/main.go
CMD ["./bin/main"]