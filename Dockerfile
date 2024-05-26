FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o test2gis cmd/main.go
CMD ["./test2gis"]