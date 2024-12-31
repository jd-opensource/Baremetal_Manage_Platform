global:
  resolve_timeout: 5s
route:
  group_by: ['alertname']
  group_wait: 0s
  group_interval: 60m
  repeat_interval: 10s
  receiver: 'bmpAlertReceiver'
  routes:
  - receiver: 'bmpAlertReceiver'
    group_interval: 5m
    match:
      noticePeriodLabel: NoticePeriod-5m
  - receiver: 'bmpAlertReceiver'
    group_interval: 10m
    match:
      noticePeriodLabel: NoticePeriod-10m
  - receiver: 'bmpAlertReceiver'
    group_interval: 15m
    match:
      noticePeriodLabel: NoticePeriod-15m
  - receiver: 'bmpAlertReceiver'
    group_interval: 30m
    match:
      noticePeriodLabel: NoticePeriod-30m
  - receiver: 'bmpAlertReceiver'
    group_interval: 60m
    match:
      noticePeriodLabel: NoticePeriod-60m
  - receiver: 'bmpAlertReceiver'
    group_interval: 180m
    match:
      noticePeriodLabel: NoticePeriod-180m
  - receiver: 'bmpAlertReceiver'
    group_interval: 720m
    match:
      noticePeriodLabel: NoticePeriod-720m
  - receiver: 'bmpAlertReceiver'
    group_interval: 1440m
    match:
      noticePeriodLabel: NoticePeriod-1440m
receivers:
  - name: 'bmpAlertReceiver'
    webhook_configs:
      - url: 'http://${BMP_PRONOEA_HOST}:${BMP_PRONOEA_PORT}/api/alert/receiver'