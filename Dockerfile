# Use the official Golang image as a base image
FROM golang:1.21-alpine3.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files to the container
COPY go.mod go.sum ./

# Use Go official module proxy
ENV GOPROXY=https://proxy.golang.org

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a lightweight alpine image for the final image
FROM alpine:3.18.3  

# Set the working directory
WORKDIR /root/

# Install bash (required for wait-for-it)
RUN apk add --no-cache bash

# Download wait-for-it
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Use wait-for-it to wait for the PostgreSQL service and then run the application
CMD ["/wait-for-it.sh", "nomess-postgres-db:5432", "--", "./main"]