FROM golang:alpine as BUILDER

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build .


FROM alpine as RUNNER

WORKDIR /app

COPY --from=BUILDER /app/prom-app .

RUN chmod +x prom-app

CMD ["./prom-app"]