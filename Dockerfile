FROM golang:1.23.1-bookworm AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api.go

FROM scratch AS run
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
ENTRYPOINT  ["./api"]
