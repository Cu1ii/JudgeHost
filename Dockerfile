FROM gleaming/golang1.9.3:env
MAINTAINER cu1

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV CGO_ENABLED 0

WORKDIR $GOPATH/src/judge-host

ADD . $GOPATH/src/judge-host

RUN mkdir build && cd build && cmake ../JudgerCore && make && make install && cd .. && go mod tidy

RUN cd $GOPATH/src/judge-host && go build .

ENTRYPOINT ["./judge-host"]