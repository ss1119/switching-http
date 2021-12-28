FROM ubuntu:18.04
WORKDIR /work
RUN apt-get update && \
    apt-get install software-properties-common -y && \
    add-apt-repository ppa:longsleep/golang-backports && \
    apt-get update && \
    apt-get install golang-go iproute2 -y 
RUN apt -y install iperf3
RUN apt-get install sudo
RUN apt-get update && apt-get install -y \
curl
COPY ./ /work
RUN go mod tidy
COPY letsEncrypt.pem /usr/local/share/ca-certificates/letsEncrypt.crt
RUN apt-get update -y && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/* && update-ca-certificates
# COPY /etc/ssl/certs/ca-certificates.crt certs.crt
# VOLUME /certs.crt:/etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
EXPOSE 80
EXPOSE 443
CMD ["go", "run", "main.go"]
