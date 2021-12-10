FROM ubuntu:18.04
WORKDIR /work
RUN apt-get update && \
    apt-get install software-properties-common -y && \
    add-apt-repository ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install golang-go iproute2 -y 
RUN apt -y install iperf3
COPY ./ /work
RUN go mod tidy
COPY localhost19.ml/localhost19.pem /usr/local/share/ca-certificates/localhost19.crt
RUN apt-get update -y && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/* && update-ca-certificates
CMD ["go", "run", "main.go"]