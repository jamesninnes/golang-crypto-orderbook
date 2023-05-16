# Use the official Golang image as the base image
FROM golang:1.17-alpine

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["/app/main"]
