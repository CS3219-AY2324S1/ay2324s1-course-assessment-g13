# syntax=docker/dockerfile:1

# Retrieves a base Go image
FROM golang:1.20

# Sets a project directory to serve as destination for COPY
WORKDIR /app

# Download Go modules by copying our local go.mod and go.sum into the image
COPY go.mod go.sum ./
# Downalods dependencies
RUN go mod download

# Copy the source code. Note the slash at the end as Docker is sensitive to trailing slash
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /question-service

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/question-service"]