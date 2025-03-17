FROM golang:1.19

WORKDIR /app/bridge-funds-transporter

RUN go install github.com/cosmtrek/air@v1.49

COPY bridge-funds-transporter/go.mod /app/bridge-funds-transporter
COPY bridge-funds-transporter/.air.toml /app/bridge-funds-transporter
COPY bridge-funds-transporter/.env /app/bridge-funds-transporter

RUN go mod download

COPY bridge-funds-transporter /app/bridge-funds-transporter

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]