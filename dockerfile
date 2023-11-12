# Use the official Golang image as the parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install SQLite3
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Build the Go app
RUN go build -o server 

# Expose port 8080 for the container
EXPOSE 8080

# Start the application
ENTRYPOINT [ "/app/server" ] 
