FROM golang:1.10

RUN mkdir /go/src/go-test-project
WORKDIR /go/src/go-test-project

COPY . .

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build -o /app/main .
CMD ["/app/main"]