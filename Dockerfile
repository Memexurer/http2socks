FROM alpine:latest
WORKDIR /opt/app

COPY build/httptosocks /opt/app/httptosocks

ENTRYPOINT ["/opt/app/httptosocks"]