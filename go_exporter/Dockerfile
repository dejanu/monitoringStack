from golang:1.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY index.html ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-prom-exporter

CMD ["/go-prom-exporter"]
