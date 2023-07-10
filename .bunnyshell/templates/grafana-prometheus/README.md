# Template overview

This template can be used to create a new environment for a monitoring stack using based on Grafana, Prometheus and a Exporter.

The template provides the Bunnyshell configuration for a monitoring stack composed of 3 Components (Grafana + Prometheus + Golang exporter).

You can extend the template by further adding Components, be them more APIs or other services, such as queues, caches, block storage etc.

&nbsp;

# Environment overview

This environment is comprised of 2 components:

- `Grafana` web-app for visualization
- `Prometheus` Monitoring system and time series database
- `scrapeapp` Golang exporter

The Grafana creds are declared as component variables.