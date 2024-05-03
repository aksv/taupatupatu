# syntax=docker/dockerfile:1.0

#######################################
# STAGE 1
#######################################

FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/messaging-app ./cmd/messaging/main.go

#######################################
# STAGE 2
#######################################

FROM scratch AS final

WORKDIR /

COPY --from=build /bin/messaging-app /messaging-app

EXPOSE 8084

ENTRYPOINT ["/messaging-app"]