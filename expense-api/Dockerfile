FROM golang:1.22.4

COPY . /app
COPY ./go-common /app

WORKDIR /app/expense-api

RUN go build -o ./bin/expense-api.exe ./cmd/server
RUN chmod +x ./bin/expense-api.exe

ENTRYPOINT [ "./bin/expense-api.exe" ]