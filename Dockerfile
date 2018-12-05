FROM golang:1.11-alpine as builder
RUN apk add --update --no-cache git
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .

FROM alpine:3.8
RUN addgroup -g 1000 corby \
    && adduser -u 1000 -G corby -s /bin/sh -D corby
COPY --from=builder /app/main /app
USER 1000
EXPOSE 8080
CMD ["/app"]
