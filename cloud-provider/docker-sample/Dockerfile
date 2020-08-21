FROM golang

WORKDIR /go/src/app
COPY *.go .
RUN go install .

CMD ["app"]