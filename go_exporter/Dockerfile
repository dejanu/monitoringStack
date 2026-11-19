FROM golang:1.20 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY index.html ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-prom-exporter .

FROM scratch
COPY --from=build /app/go-prom-exporter /go-prom-exporter
COPY --from=build /app/index.html /index.html

EXPOSE 5000
CMD ["/go-prom-exporter"]
