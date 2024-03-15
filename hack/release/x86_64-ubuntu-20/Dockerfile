FROM ubuntu:20.04

RUN apt-get update

RUN apt-get install -y ca-certificates curl git make build-essential

COPY --from=golang:1.21 /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir /workspace

RUN mkdir /workspace/_output

RUN cd /workspace && git clone https://github.com/OKESTRO-AIDevOps/nkia.git 

RUN go clean -cache && go clean -modcache

RUN cd /workspace/nkia && go mod tidy && go get all

RUN cd /workspace/nkia && make release

RUN cp /workspace/nkia/nkia.tgz /workspace/_output

CMD ["/bin/bash", "-c", "/bin/cp -Rf /workspace/nkia/nkia.tgz /workspace/_output; /bin/cp -Rf /workspace/nkia/lib.tgz /workspace/_output; echo SUCCESS."]