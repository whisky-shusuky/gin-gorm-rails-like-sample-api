FROM golang:1.16

ENV TZ='Asia/Tokyo'
RUN apt-get -y update && \
  go get -u github.com/cespare/reflex && \
  ln -s /var/src/gin-gorm-rails-like-sample-api /go/src/gin-gorm-rails-like-sample-api

CMD ["/bin/bash"]