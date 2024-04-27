FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

RUN go build -o main .

# Command to run the executable
CMD ["./main"]