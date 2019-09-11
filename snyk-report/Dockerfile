FROM golang:latest

COPY entrypoint.sh /go/entrypoint.sh
COPY main.go /snyk-report/main.go
RUN chmod +x /go/entrypoint.sh

ENTRYPOINT ["/go/entrypoint.sh"]