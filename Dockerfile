FROM golang:latest 

ENV GOPATH /go

RUN mkdir -p /go/src/github.com/droslean/govue-quiz
ADD . /go/src/github.com/droslean/govue-quiz/
WORKDIR /go/src/github.com/droslean/govue-quiz/

RUN go build -o main . 
CMD ["/go/src/github.com/droslean/govue-quiz/main"]
