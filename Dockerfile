# build app
FROM golang:alpine3.20 as builder

WORKDIR /jornada-milhas

COPY . .

RUN go install ./cmd/jornada-milhas

# deploy binary
FROM alpine:latest

WORKDIR /

COPY --from=builder /go/bin/jornada-milhas /jornada-milhas
RUN chmod 700 /jornada-milhas

EXPOSE 3333

ENTRYPOINT /jornada-milhas