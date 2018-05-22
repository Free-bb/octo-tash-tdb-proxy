FROM golang:1.10
ENV LD_LIBRARY_PATH /usr/local/lib

# TrailDB Steps
RUN apt-get update
RUN apt-get install -y libarchive-dev libjudy-dev pkg-config python
WORKDIR /opt/traildb
ADD https://github.com/traildb/traildb/archive/0.6.tar.gz .
RUN tar xfz 0.6.tar.gz
WORKDIR /opt/traildb/traildb-0.6
RUN ./waf configure
RUN ./waf install

# FROM golang:1.10

# LABEL maintainer="Thomas P <thomas@scullwm.com>"

RUN mkdir -p /app
RUN mkdir -p /var/log/tdb
RUN touch /var/log/tdb/main.log

WORKDIR /app

ADD main.go /app
ADD routing.yaml /app
RUN go get -u -v gopkg.in/yaml.v2
RUN go get -u -v github.com/traildb/traildb-go
RUN go get -u -v golang.org/x/blog/content/context/userip
RUN go build ./main.go

# CMD ["./main"]

EXPOSE 7000
