# This Dockerfile is not optimized on purpose, since it's
# likely to evolve, let's keep the building packages in
# the cache.


FROM debian:jessie
MAINTAINER s. rannou <mxs@sbrk.org>


ENV DEBIAN_FRONTEND noninteractive
RUN mkdir /go /app
ENV GOPATH /go


RUN apt-get update
RUN apt-get install -q -y golang git mercurial
RUN go get github.com/gin-gonic/gin
RUN go get github.com/dancannon/gorethink
ADD src /src
RUN ( cd /src && go build -o /app/ogipa )
ADD bin/ogipa.json /app/ogipa.json


CMD (cd /app && ./ogipa -c ogipa.json )
