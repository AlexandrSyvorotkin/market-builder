FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -v .
CMD ["./server"]
EXPOSE 5000
