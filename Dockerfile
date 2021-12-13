FROM golang:1.16-alpine

WORKDIR /app

# Download dependencies first
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Then, copy source code
COPY . .

# Build
RUN go build -o "./reas"

# Run
EXPOSE 5000
CMD ["./reas"]