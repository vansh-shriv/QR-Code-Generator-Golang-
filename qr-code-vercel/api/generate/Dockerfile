FROM golang:1.22 as build

WORKDIR /app
COPY handler.go .

RUN go build -o app handler.go

FROM debian:bullseye-slim

COPY --from=build /app/app /app
CMD ["/app/app"]
