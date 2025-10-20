FROM golang:1.25-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o app main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 8080
CMD [ "/app/app" ]