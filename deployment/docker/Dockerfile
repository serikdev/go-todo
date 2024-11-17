# Using the basic image on Go
FROM golang:1.22.8

# Installing the working directory
WORKDIR /app

# Copy go.mod and go.sum to load dependencies
COPY go.mod go.sum ./ 
RUN go mod download

# Copy the entire project
COPY . .

# Compiling the application
RUN go build -o main ./cmd

# Launching the application
CMD ["/app/main"]

# Specify that the container uses port 8080
EXPOSE 8080
