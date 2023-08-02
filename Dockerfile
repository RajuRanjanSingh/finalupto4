FROM golang:latest
WORKDIR /app
COPY . /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 4747