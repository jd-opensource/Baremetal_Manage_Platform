scrape_configs:
  - job_name: 'pushgateway'
    scrape_interval: 60s
    honor_labels: true
    static_configs:
      - targets: ['${BMP_PUSHGATEWAY_HOST}:${BMP_PUSHGATEWAY_PORT}']
alerting:
  alertmanagers:
    - static_configs:
       - targets: ['${BMP_ALERTMANAGER_HOST}:${BMP_ALERTMANAGER_PORT}']

rule_files:
  - /var/lib/prometheus/conf/rules/*.yml