FROM alpine
RUN apk add --no-cache wget ffmpeg git go npm openssh

VOLUME [ "/usr/local/maidnana/go-cqhttp", "/usr/local/maidnana/MaidNanaGo" ]
WORKDIR /usr/local/maidnana/go-cqhttp

RUN wget https://github.com/Mrs4s/go-cqhttp/releases/download/v1.0.0-rc3/go-cqhttp_linux_amd64.tar.gz && tar -xvf go-cqhttp_linux_amd64.tar.gz

WORKDIR /usr/local/maidnana/src

RUN git clone https://github.com/nanoyou/MaidNanaGo && git clone https://github.com/nanoyou/MaidNanaFrontEnd

WORKDIR /usr/local/maidnana/src/MaidNanaFrontEnd
RUN npm install && npm run build && mv dist ../MaidNanaGo/MaidNanaFrontEnd/

WORKDIR /usr/local/maidnana/src/MaidNanaGo
RUN export GOPATH=/usr/lib/go && go install github.com/swaggo/swag/cmd/swag@latest
RUN /usr/lib/go/bin/swag init && go build -o /usr/local/maidnana/MaidNanaGo/MaidNanaGo

CMD cd /usr/local/maidnana/go-cqhttp && echo 02 > ./go-cqhttp & cd /usr/local/maidnana/MaidNanaGo && ./MaidNanaGo
