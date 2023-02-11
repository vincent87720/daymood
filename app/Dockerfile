FROM golang:1.19 AS build
WORKDIR /app
ENV GO111MODULE=on \
    CGO_ENABLED=0
COPY app /app
RUN go build -o /app/bin/daymood /app/cmd/daymood/main.go


FROM alpine:latest AS release
ENV APP_MODE=PROD \
    APP_HOST= \
    APP_PORT=8000 \
    DB_HOSTNAME=daymood-database \
    DB_DATABASE=daymood \
    DB_USERNAME=daymood \
    DB_PASSWORD=daymood \
    EZSTORE_USERNAME= \
    EZSTORE_PASSWORD= \
    SESSION_SECRET=daymood
COPY --from=build /app/bin/daymood /app/daymood
COPY --from=build /app/systemConfigs.json /app/systemConfigs.json
COPY --from=build /app/settings.yaml /app/settings.yaml
ENTRYPOINT ["/app/daymood"]