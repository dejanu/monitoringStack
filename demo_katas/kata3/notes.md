* Prometheus + Alert Rules

```bash
docker-compose up -d

# trigger the alert
docker stop node_exporter

# up == 0 has no label filter, so it evaluates for every target Prometheus scrapes
# -  not just node_exporter
```