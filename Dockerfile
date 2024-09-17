FROM golang:1.23 as builder

COPY ./docker-entrypoint-initdb.d/init.sql ./docker-entrypoint-initdb.d/
RUN chmod 644 ./docker-entrypoint-initdb.d/init.sql

COPY .env ./

WORKDIR /backend

COPY /backend/go.mod /backend/go.sum ./
RUN go mod download

COPY ./backend ./

RUN go build -v -o /backend/main ./app/main.go

FROM golang:1.23

WORKDIR /backend
COPY --from=builder /backend/main ./app/main

RUN chmod +X ./app/main

EXPOSE 8080

CMD ["./app/main"]