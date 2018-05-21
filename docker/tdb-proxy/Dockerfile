FROM golang:1.10

LABEL maintainer="Thomas P <thomas@scullwm.com>"

RUN mkdir -p /app
RUN mkdir -p /var/log/tdb
RUN touch /var/log/tdb/main.log

WORKDIR /app

ADD main.go /app
ADD routing.yaml /app
RUN go get -u -v gopkg.in/yaml.v2
RUN go build ./main.go

CMD ["./main"]

EXPOSE 7000
