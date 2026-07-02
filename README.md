# monitoringStack
Monitoring stack comprised of Grafana/Prometheus/Golang

![Monitoring Stack](src/stack.PNG)

## Local setup

* Spin up standalone [Go web-app/exporter](https://hub.docker.com/repository/docker/dejanualex/go_promexporter/general):
```
docker run -p 5000:5000 dejanualex/go_promexporter:1.1
```
* Spin up the stack
```bash
# check config
docker-compose -f docker-compose.yml config
# start in foreground
docker-compose -f docker-compose.yml up --remove-orphans
# stop services
docker-compose -f docker-compose.yml down
```
* Prometheus should be available : http://127.0.0.1:9090/
* Grafana should be available: http://127.0.0.1:3000/login
* Blackbox Go exporter endpoints: http://127.0.0.1:5000/ and http://127.0.0.1:5000/metrics .

    - Check `promhttp_metric_handler_requests_total` standard Prometheus Go client counter metric that tracks the total number of HTTP requests made to an application's `/metric`
    - Check series `promhttp_metric_handler_requests_total{instance="goexporter:5000"}` and `docker restart localgoexporter` to reset the counter

## Bunnyshell setup

![Monitoring Stack](src/services.png)