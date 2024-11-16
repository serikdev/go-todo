
FROM golang:1.22.8


WORKDIR /app


COPY go.mod go.sum ./ 
RUN go mod download

COPY . .


RUN go build -o main ./cmd


CMD ["/app/main"]


EXPOSE 8080
