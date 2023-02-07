FROM golang:1.19.5-alpine
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o ./...
CMD ["go", "run", "main.go"]
