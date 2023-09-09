FROM golang:1.20

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install air
RUN go install github.com/cosmtrek/air@latest

# Install air binary
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# Sets alias as /bin/air
RUN alias air='$(go env GOPATH)/bin/air'

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air"]