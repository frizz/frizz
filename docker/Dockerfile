FROM ubuntu:latest

RUN echo "LAST UPDATE 5th JAN 2018"

RUN apt-get update && apt-get install -y \
  git \
  build-essential \
  libreadline-dev \
  wget \
  mercurial meld \
  curl \
  cmake \
  python-dev \
  && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y software-properties-common

############# Fish #############
RUN echo 'deb http://download.opensuse.org/repositories/shells:/fish:/release:/2/Debian_9.0/ /' > /etc/apt/sources.list.d/fish.list 
RUN apt-get update
RUN apt-get install -y --allow-unauthenticated fish

############# Golang #############
RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
RUN mkdir /go
RUN export GOPATH=/root/
ENV GOPATH /root/
RUN export PATH=$PATH:$GOPATH/bin
ENV PATH $PATH:$GOPATH/bin
RUN export PATH=$PATH:/usr/local/go/bin
ENV PATH $PATH:/usr/local/go/bin

RUN apt-get install -y nodejs
RUN apt-get install -y npm
RUN ln -s /usr/bin/nodejs /usr/bin/node

ENV NODE_PATH /usr/local/lib/node_modules

RUN npm install --global source-map-support

############# Circle CI #############
# RUN curl -o /usr/local/bin/circleci https://circle-downloads.s3.amazonaws.com/releases/build_agent_wrapper/circleci && chmod +x /usr/local/bin/circleci

############# Docker #############
# RUN apt-get install -y \
#     apt-transport-https \
#     ca-certificates \
#     curl \
#     software-properties-common
# RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
# RUN apt-key fingerprint 0EBFCD88
# RUN add-apt-repository \
#        "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
#        $(lsb_release -cs) \
#        stable"
# RUN apt-get update
# RUN apt-get install -y docker-ce

WORKDIR /root/src/github.com/gopherjs/gopherjs
