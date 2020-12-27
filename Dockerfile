FROM golang:1.14-stretch as builder
RUN mkdir -p /go/src/github.com/themoah/go-dead-drop
COPY . /go/src/github.com/themoah/go-dead-drop
WORKDIR /go/src/github.com/themoah/go-dead-drop 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-dead-drop .
FROM alpine:3
RUN apk --no-cache add ca-certificates
LABEL maintainer="@themoah" 
LABEL version="0.1"
LABEL description="go-dead-drop"
RUN adduser -S -D -H -h /app appuser
WORKDIR /app
USER appuser
COPY --from=builder /go/src/github.com/themoah/go-dead-drop/go-dead-drop /app/
EXPOSE 8080
CMD ["./go-dead-drop"]