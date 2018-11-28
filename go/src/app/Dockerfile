FROM golang

ADD . /go/src/cambio-api

RUN go get golang.org/x/net/html
RUN go get github.com/lib/pq
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/stretchr/testify
RUN go get github.com/jasonlvhit/gocron

RUN go install cambio-api

ENTRYPOINT /go/bin/cambio-api

EXPOSE 8000
