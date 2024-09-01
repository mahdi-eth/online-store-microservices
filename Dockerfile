# Dockerfile for generating Go gRPC code
FROM golang:1.20-buster

# Install the necessary tools
RUN apt-get update && apt-get install -y \
    git protobuf-compiler make \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set the working directory
WORKDIR /app

# Copy the proto files into the container
COPY proto/ /app/proto/

# Run the protoc command
CMD ["make", "proto"]
