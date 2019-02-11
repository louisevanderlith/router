FROM alpine:latest as builder

COPY router .
COPY conf conf

ENTRYPOINT [ "./router" ]
