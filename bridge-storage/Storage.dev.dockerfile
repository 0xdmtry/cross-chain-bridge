FROM golang:1.19

WORKDIR /app/bridge-storage

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-storage/go.mod /app/bridge-storage
COPY bridge-storage/.air.toml /app/bridge-storage
COPY bridge-storage/.env /app/bridge-storage

RUN go mod download

COPY bridge-storage /app/bridge-storage

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]