FROM golang

ENV GOBIN /go/bin

RUN curl https://glide.sh/get | sh

#RUN go get -u github.com/flowup/owl/cmd/owl

# Test watcher
RUN go get github.com/smartystreets/goconvey

# goconvey port
EXPOSE 8080