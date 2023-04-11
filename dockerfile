## Build
FROM golang:latest AS build

WORKDIR /app

COPY go.mod ./ 
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o gin-mongo .

## Deploy
FROM scratch

COPY --from=build /app/gin-mongo /opt/gin-mongo
EXPOSE 8085

CMD ["/opt/gin-mongo"]