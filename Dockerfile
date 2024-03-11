FROM golang:1.21-alpine as builder

RUN apk --no-cache add -U tzdata
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /querofreela ./server

FROM scratch as production
LABEL authors="schivei"
WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /querofreela .

EXPOSE 8000

CMD ["/app/querofreela"]
