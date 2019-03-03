FROM alpine:latest

COPY router .
COPY conf conf

EXPOSE 8080

ENTRYPOINT [ "./router" ]
