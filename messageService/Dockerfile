FROM golang:1.7
RUN go get github.com/gorilla/mux gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson
WORKDIR /go/src/app
ADD . /go/src/app
RUN go install
CMD ["app"]
