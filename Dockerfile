FROM ubuntu:22.04

# Install golang 1.24
RUN apt-get update && apt-get install -y wget tar && \
    wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz && \
    rm go1.24.0.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
# Install git
RUN apt-get update && apt-get install -y git

