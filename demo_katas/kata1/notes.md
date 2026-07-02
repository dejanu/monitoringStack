
```bash

# start prom instance
docker run --rm  --name prom -d -p 9091:9090 prom/prometheus

# main prometheus that scrapes other prometheus (called prom)
docker run --rm -d --name prometheus -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml   prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-lifecycle

# check the prometheus configuration file
docker exec -it prometheus sh
cat /etc/prometheus/prometheus.yml

# reload prometheus to pick-up config changes
curl -X POST http://localhost:9090/-/reload

curl http://localhost:9091/-/healthy
curl http://localhost:9091/-/ready
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

# check DNS resolution
docker exec prometheus nslookup prom

docker-compose -f docker-compose.yaml down

## ALTERNATIVE create ur own network
docker network create demo
docker run --network demo --name prom ...
```