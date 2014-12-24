# This Dockerfile is not optimized on purpose, since it's
# likely to evolve, let's keep the building packages in
# the cache.


FROM debian:jessie
MAINTAINER s. rannou <mxs@sbrk.org>


ENV DEBIAN_FRONTEND noninteractive
RUN mkdir /go
ENV GOPATH /go


ADD src /src
ADD bin /app
RUN apt-get update
RUN apt-get install -q -y golang git
RUN go get github.com/gin-gonic/gin
RUN ( cd /src && go build -o /app/ogipa )


CMD (cd /app && ./ogipa -c ogipa.json )
