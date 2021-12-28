FROM ubuntu:18.04 
WORKDIR /work
RUN apt-get update && \
    apt-get install iproute2 iputils-ping -y
RUN apt -y install iperf3
RUN apt-get update && apt-get install -y \
curl
RUN apt-get update && apt-get install -y wget
RUN apt-get update && apt-get install -y gnupg2
RUN apt-get install sudo

# RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
# RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
# RUN rm google-chrome-stable_current_amd64.deb

COPY letsEncrypt.pem /usr/local/share/ca-certificates/letsEncrypt.crt
RUN update-ca-certificates

# RUN wget -q -O – https://dl.google.com/linux/linux_signing_key.pub | apt-key add && \
#     echo 'deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main' | tee /etc/apt/sources.list.d/google-chrome.list && \
#     apt-get update && \
#     apt-get install -y google-chrome-stable

# CMD curl https://localhost19.ml:8080/