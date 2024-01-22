# .devcontainer/Dockerfile
FROM golang:latest

# Install fastText dependencies
RUN apt-get update && \
    apt-get install -y build-essential libboost-all-dev

# Copy the entire project into the container
COPY . /workspace

# Set the working directory
WORKDIR /workspace
