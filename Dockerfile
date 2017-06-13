FROM golang

ADD . /go/src/lowes.com/jxriddle/store-services
RUN cd /go/src/lowes.com/jxriddle/store-services; go get
RUN go install lowes.com/jxriddle/store-services

EXPOSE 8080

ENTRYPOINT /go/bin/store-services
