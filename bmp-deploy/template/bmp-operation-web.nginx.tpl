log_format  main1  '$time_local | $remote_addr | $upstream_addr | $http_host | $status | $body_bytes_sent | $upstream_response_time | $request_time | $http_host | $request | $http_referer';

server {
    listen 8081;
    server_name bmp-operation.bmp.local;
    charset utf-8;
    access_log  /var/log/nginx/access.log  main1;

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root html;
    }

    root /home/bmp/bmp-operation-web;
    index index.html;

    # compression-webpack-plugin 配置
    #gzip on;
    #gzip_min_length 1k;
    #gzip_comp_level 9;
    #gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript applicationx-httpd-php image/jpeg image/gif image/png;
    #gzip_vary on;
    # 配置禁用 gzip 条件，支持正则，此处表示 ie6 及以下不启用 gzip（因为ie低版本不支持）
    #gzip_disable "MSIE [1-6]\.";

    location /operation-web/ {
        client_max_body_size 51200M;
        resolver 127.0.0.11 ipv6=off valid=2s;
        set $bmp_operation_api "bmp-operation-api:8799";
        rewrite ^/operation-web/(.*) /$1 break;
        proxy_pass http://$bmp_operation_api;      
    }
    
    location /oob-alert/ {
        client_max_body_size 51200M;
        resolver 127.0.0.11 ipv6=off valid=2s;
        set $bmp_oob_alert "bmp-oob-alert:8804";
        rewrite ^/oob-alert/(.*) /$1 break;
        proxy_pass http://$bmp_oob_alert;      
    }
    
    location / {
        try_files $uri $uri/ @router;
        index index.html;
    }

    location @router {
        rewrite ^.*$ /index.html last;
    }
    location ~.*\.(html)$ {
	add_header Cache-Control no-cache;
    }

}
