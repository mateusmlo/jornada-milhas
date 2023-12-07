FROM golang
ADD . /go/src/github.com/mateusmlo/jornada-milhas
WORKDIR /go/src/github.com/mateusmlo/jornada-milhas

# SERVER
ARG SERVER_PORT=3333
ARG APP_MODE=prod

# DB
ARG DB_HOST=db
ARG DB_USER=jornada-admin
ARG DB_PASSWORD
ARG DB_NAME=jornadas-db
ARG DB_PORT=5432

# JWT
ARG TOKEN_TTL=2
ARG JWT_SECRET

# REDIS
ARG REDIS_HOST=cache
ARG REDIS_PORT=6379
ARG REFRESH_TTL=864000

RUN go install ./cmd/jornada-milhas

ENTRYPOINT /go/bin/jornada-milhas
EXPOSE 3333