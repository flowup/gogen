FROM golang:1.7.4

# get testing service
RUN go get github.com/smartystreets/goconvey

# get auto reloading service
RUN go get github.com/codegangsta/gin

# expose goconvey port
EXPOSE 8080

# workdir of the project
WORKDIR /go/src/github.com/flowup/gogen
