FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | apt-key add -
RUN echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list
RUN apt-get update

RUN DEBIAN_FRONTEND=noninteractive apt-get -y install lsb-release wget devscripts debhelper debconf pkg-config curl make git libkrb5-dev git mongodb-mongosh
RUN wget https://golang.org/dl/go1.19.linux-arm64.tar.gz -O /tmp/golang1.19.tar.gz
RUN tar --transform=s,go,go1.19, -zxf /tmp/golang1.19.tar.gz
RUN mv go1.19 /usr/local/
RUN ln -s /usr/local/go1.19/bin/go /usr/bin/go
