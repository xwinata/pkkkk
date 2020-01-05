FROM golang:alpine

ARG Root="pkkkk"

ADD . "${ROOT}"
WORKDIR "${ROOT}"

RUN apk add --update git gcc libc-dev tzdata;
#  wget gcc libc-dev make openssl py-pip;

ENV TZ=Asia/Jakarta

RUN go get -u github.com/golang/dep/cmd/dep

CMD dep ensure -v \
    && go build -v -o "${ROOT}" \
    # run app mode
    && "${APPNAME}" run ;

EXPOSE 8000