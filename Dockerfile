FROM golang:latest AS builder

WORKDIR /app

COPY ./ ./

RUN go mod download


#после -o название бинарника, дальше путь до точки входа
RUN go build -o bin ./cmd/main.go

CMD [ "/app/bin" ]