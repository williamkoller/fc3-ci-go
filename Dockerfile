FROM golang:1.15

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]