FROM golang

ADD . /go/src/wscraper

RUN go get golang.org/x/net/html
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm

RUN go install wscraper

ENTRYPOINT /go/bin/wscraper

EXPOSE 8000


