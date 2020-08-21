from golang:latest

RUN mkdir /application
ADD . /application
WORKDIR /application
RUN go build -o application .

CMD ["/application/application"]
