# Use the official Golang image as the base image
FROM golang:1.22 as builder

# Set the working directory within the builder container
WORKDIR /build

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the builder container
COPY . .

# Print contents of /build directory for debugging
RUN ls -l /build

# Build the Go app for Linux (amd64) with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-list .

# Create a new stage for the final application image (based on Alpine Linux)
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /build/todo-list ./todo-list

# Copy your configuration files
COPY --from=builder /build/app.env ./app.env

# Copy your migration files
COPY --from=builder /build/db/migrations ./db/migrations

# Expose port 8080 if your application needs it
EXPOSE 8080

# Command to run your application
CMD ["./todo-list"]
