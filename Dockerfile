# Use official Golang base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Disable Go modules for this simple app
ENV GO111MODULE=off

# Build the Go app
RUN go build -o main .

# Expose the app's port
EXPOSE 8080

# Command to run the app
CMD ["./main"]

