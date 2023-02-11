# Daymood APP

## Environment
- go1.19.4 darwin/arm64

## Development

### Environment variables
Create the `.env` file in the app directory
```
APP_MODE=DEV
APP_HOST=
APP_PORT=8000
DB_HOSTNAME=database #docker-compose service name or database IP addr
DB_DATABASE=daymood
DB_USERNAME=daymood
DB_PASSWORD=daymood
EZSTORE_USERNAME=
EZSTORE_PASSWORD=
SESSION_SECRET=daymood
```

### Compiles for development
```
make run
```

### Compiles for development using docker
```
make dockerrun
```

### Start using services
Use any tool that can send HTTP requests to the API (e.g. Postman) to send requests and parameters to the api located at 8000 port.

## Build
### Build binary
```
make build
```
### Build docker images
[DaymoodReadme#Build](../README.md#建置)