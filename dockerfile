## Build
FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./ 
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

## Deploy
FROM scratch

WORKDIR /opt/sample-opentracer
COPY --from=build /app/main /opt/gin-mongo/main
EXPOSE 8080

CMD ["./main"]