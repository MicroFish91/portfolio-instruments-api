# syntax=docker/dockerfile:1
FROM golang:1.21.5-alpine

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/portfolio-instruments ./cmd/api

# Port
EXPOSE 3000

# Run
CMD ["/bin/portfolio-instruments"]