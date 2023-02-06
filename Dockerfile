FROM golang:1.19.5-alpine
COPY go.mod .
RUN go mod download
COPY . .
RUN go build
CMD ["go", "run", "main.go"]