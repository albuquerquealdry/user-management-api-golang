FROM golang:1.23.3
WORKDIR /go/src/app
COPY . . 
EXPOSE 8000
RUN go build -o main src/cmd/main.go
CMD ["./main"]