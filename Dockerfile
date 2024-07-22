FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o golang-practice main.go
EXPOSE 8080
CMD ["/app/golang-practice"]
