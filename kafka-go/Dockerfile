FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY *.go go.mod .
RUN go mod tidy
#RUN go get yaml-hostname-updater
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yaml-hostname-updater .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /data
WORKDIR /data
COPY --from=builder /app/yaml-hostname-updater /root/
RUN chmod +x /root/yaml-hostname-updater
ENTRYPOINT ["/root/yaml-hostname-updater"]
