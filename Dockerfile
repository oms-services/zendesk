FROM golang

RUN go get github.com/gorilla/mux

RUN go get -v github.com/aws/aws-sdk-go/aws

RUN go get -v github.com/MEDIGO/go-zendesk/zendesk

WORKDIR /go/src/github.com/oms-services/zendesk

ADD . /go/src/github.com/oms-services/zendesk

RUN go install github.com/oms-services/zendesk

ENTRYPOINT zendesk

EXPOSE 3000