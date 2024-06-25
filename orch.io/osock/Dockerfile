FROM ubuntu:20.04

RUN mkdir /osock

WORKDIR /osock

RUN apt-get update 

RUN apt-get install -y ca-certificates

COPY --from=golang:1.21 /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

COPY . .


RUN go clean -modcache

RUN go mod init orch.io/osock

RUN go get github.com/OKESTRO-AIDevOps/nkia@main


RUN go mod tidy

RUN go build -o osock .

CMD ["/osock/osock"]