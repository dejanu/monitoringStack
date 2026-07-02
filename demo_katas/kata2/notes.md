

```bash
# create network
docker network create promnet

# start node exporter
docker run --network promnet --name node_exporter -d \
  -p 9100:9100 \
  prom/node-exporter

# start prometheus
docker run  --network promnet --rm -d --name prometheus -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml   prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-lifecycle

# cleanup
docker stop prometheus node_exporter
docker network rm promnet


# start setup as docker-compose
docker-compose -f docker-compose.yaml up -d
docker-compose -f docker-compose.yaml down
```