FROM ubuntu:20.04

RUN apt-get update

RUN apt-get install -y ca-certificates curl

COPY --from=golang:1.21 /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir /workspace

WORKDIR /workspace

COPY . .

WORKDIR /workspace/srv

RUN go build -o runo .

CMD ["/workspace/srv/runo"]