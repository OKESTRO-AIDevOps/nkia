FROM ubuntu:20.04

RUN apt-get update

RUN apt-get install -y ca-certificates curl git make build-essential

COPY --from=golang:1.20 /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir /workspace

RUN mkdir /workspace/_output

RUN cd /workspace && git clone https://github.com/OKESTRO-AIDevOps/nkia.git 

RUN go clean -cache && go clean -modcache

RUN cd /workspace/nkia && go get all

RUN cd /workspace/nkia && make release

RUN cp /workspace/nkia/nkia.tgz /workspace/_output

CMD ["echo","SUCCESS."]