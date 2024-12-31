module(load="imfile")
input(type="imfile" File="{{log_path}}" Tag="bmp-agent")
if $programname == 'bmp-agent' then @{{rsyslog_host}}:{{rsyslog_port}}