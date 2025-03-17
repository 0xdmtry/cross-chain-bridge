FROM golang:1.19

WORKDIR /app/bridge-accounts-creator

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-accounts-creator/go.mod /app/bridge-accounts-creator
COPY bridge-accounts-creator/.air.toml /app/bridge-accounts-creator
COPY bridge-accounts-creator/.env /app/bridge-accounts-creator

RUN go mod download

COPY bridge-accounts-creator /app/bridge-accounts-creator

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]