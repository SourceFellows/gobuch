FROM golang:latest

COPY main.go /app/main.go
WORKDIR /app
RUN go build main.go

CMD [ "/app/main"]