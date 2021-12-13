FROM golang:1.16-alpine
LABEL org.opencontainers.image.source="https://github.com/gforien/reddit-assignment"

WORKDIR /app

# Download dependencies first
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Then, copy source code
COPY . .

# Build
RUN go build -o "./reas"

# Remove source code
RUN rm *.go

# Run
EXPOSE 5000
CMD ["./reas"]
