FROM 1.23.3-nanoserver-1809 AS build
WORKDIR /app
COPY . .
RUN go mod donwload
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api.go

FROM scratch AS run
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
ENTRYPOINT  ["./api"]
