FROM golang:1.12.0

RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go get go.mongodb.org/mongo-driver; exit 0
RUN go build -o moist .

CMD ["/app/moist"]
