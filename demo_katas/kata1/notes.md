* Prometheus + Prometheus

```bash
# start prom instance
docker run --rm  --name prom -d -p 9091:9090 prom/prometheus

# main prometheus that scrapes other prometheus (update prometheus.yaml to add target prom)
docker run --rm -d --name prometheus -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml  prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-lifecycle

# start the same container with docker-compose file
docker-compose -f docker-compose.prometheus.yaml up -d
docker-compose -f docker-compose.prometheus.yaml down

# check the prometheus configuration file
docker exec -it prometheus sh
cat /etc/prometheus/prometheus.yml

# reload prometheus to pick-up config changes
curl -X POST http://localhost:9090/-/reload

curl http://localhost:9090/-/healthy
curl http://localhost:9090/-/ready


```

```bash

# Container ↔ Container (one container talking directly to another)
# get container private IP
docker inspect prom --format '{{json .NetworkSettings.Networks}}'

# User-defined bridges provide automatic DNS resolution between containers
https://docs.docker.com/engine/network/drivers/bridge/

# Docker also starts an embedded DNS service for it custom-network
# Creating network "demo_default" with the default driver
docker-compose -f docker-compose.yaml up -d

# check DNS resolution (works with custom network i.e. demo network)
docker exec prometheus nslookup prom

docker-compose -f docker-compose.yaml down

## ALTERNATIVE create your own network
docker network create demo
docker run --network demo --name prom ...
```