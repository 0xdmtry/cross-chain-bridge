FROM golang:1.19

WORKDIR /app/bridge-contracts-provider

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-contracts-provider/go.mod /app/bridge-contracts-provider
COPY bridge-contracts-provider/.air.toml /app/bridge-contracts-provider
COPY bridge-contracts-provider/.env /app/bridge-contracts-provider

RUN go mod download

COPY bridge-contracts-provider /app/bridge-contracts-provider

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]