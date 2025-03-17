FROM golang:1.19

WORKDIR /app/bridge-eth-deployer

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-eth-deployer/go.mod /app/bridge-eth-deployer
COPY bridge-eth-deployer/.air.toml /app/bridge-eth-deployer
COPY bridge-eth-deployer/.env /app/bridge-eth-deployer

RUN go mod download

COPY bridge-eth-deployer /app/bridge-eth-deployer

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]