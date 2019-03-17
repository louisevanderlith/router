FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM alpine:latest

COPY --from=builder /box/router .
COPY conf conf

EXPOSE 8080

ENTRYPOINT [ "./router" ]