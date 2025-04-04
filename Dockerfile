FROM golang:1.23.5 as build

WORKDIR /src

COPY go.sum go.mod ./

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app .

FROM alpine:latest

COPY --from=build /app /app

EXPOSE 8080

CMD ["/app"]