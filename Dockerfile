FROM golang:1.21.3-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /backend-app cmd/web/main.go cmd/web/middleware.go cmd/web/router.go

EXPOSE 8080

CMD [ "/backend-app" ]