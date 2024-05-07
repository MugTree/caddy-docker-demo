# Start from a minimal base image
FROM golang:alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o bin/app

# Start from a scratch (empty) image
FROM scratch

# Copy the built Go executable from the builder stage into the empty container
COPY --from=builder /app/bin/app /

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the Go executable
CMD ["/app"]