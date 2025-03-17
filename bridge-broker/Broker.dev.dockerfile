FROM golang:1.19

WORKDIR /app/bridge-broker

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-broker/go.mod /app/bridge-broker
COPY bridge-broker/.air.toml /app/bridge-broker
COPY bridge-broker/.env /app/bridge-broker

RUN go mod download

COPY bridge-broker /app/bridge-broker

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]