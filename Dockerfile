FROM golang:latest 
WORKDIR /go/src/golang-telepresence-service
RUN go get -d -v github.com/gin-gonic/gin
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/bcrypt-service/main .
CMD ["./main"]
