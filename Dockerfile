FROM golang:1.11 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/router .
COPY conf conf

EXPOSE 8080

ENTRYPOINT [ "./router" ]