FROM golang:1 as builder

WORKDIR /go/src/app
COPY *.go .
RUN CGO_ENABLED=0 go build .

####################################

FROM alpine
COPY --from=builder /go/src/app/app .
EXPOSE 8080
CMD [ "./app" ]
