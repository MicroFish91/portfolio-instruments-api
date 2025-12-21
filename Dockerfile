# syntax=docker/dockerfile:1
FROM golang:1.25.4-alpine

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . .

# Environment variables
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOFLAGS="-mod=vendor"

# Build
RUN go build -o /bin/portfolio-instruments ./cmd/api

# Port
EXPOSE 3000

# Run
CMD ["/bin/portfolio-instruments"]