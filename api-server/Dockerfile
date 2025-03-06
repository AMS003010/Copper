# Use an official Golang image as a build stage 
FROM golang:1.23-alpine AS builder  

# Set the current working directory inside the container 
WORKDIR /app  

# Copy only the necessary Go module files first (for caching dependencies) 
COPY go.mod go.sum ./  

# Download dependencies 
RUN go mod download  

# Copy the entire source code into the container 
COPY . .   

# Build the Go app binary 
RUN go build -o api-server ./main.go  

# Build the migration binary 
RUN go build -o migrate ./migrate/migrate.go  

# Use a minimal image for running the application 
FROM alpine:3.18  

# Set the current working directory inside the container 
WORKDIR /app  

# Copy the pre-built binaries from the builder stage 
COPY --from=builder /app/api-server . 
COPY --from=builder /app/migrate .  

# Expose the port the app runs on 
EXPOSE 8080  

# Command to run the migrations and then start the server 
CMD ["sh", "-c", "./migrate && ./api-server"]