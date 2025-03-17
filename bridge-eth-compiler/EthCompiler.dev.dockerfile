# Use Ubuntu 22.04 as the base
FROM ubuntu:22.04

# Update Ubuntu Software repository
RUN apt update

# Install necessary tools including wget, curl, and Node.js
RUN apt install -y software-properties-common curl wget gnupg nodejs npm

# Install solc using npm
RUN add-apt-repository ppa:ethereum/ethereum
RUN apt update
RUN apt install solc

# Download Go 1.19 - replace the URL with the latest one from the Go website
RUN wget https://golang.org/dl/go1.19.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
RUN rm go1.19.linux-amd64.tar.gz

# Set up the Go environment
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Set a working directory
WORKDIR /app/bridge-eth-compiler

# Copy the Go modules and install dependencies
COPY bridge-eth-compiler/go.mod /app/bridge-eth-compiler
COPY bridge-eth-compiler/.air.toml /app/bridge-eth-compiler
COPY bridge-eth-compiler/.env /app/bridge-eth-compiler

RUN go mod download

COPY bridge-eth-compiler /app/bridge-eth-compiler

# Install Air for live reloading
#RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/cosmtrek/air@v1.49

# Clean up cache data and remove unnecessary packages
RUN apt clean && rm -rf /var/lib/apt/lists/*

EXPOSE 8000

# The default command to run when starting the container
CMD ["air", "-c", ".air.toml"]
