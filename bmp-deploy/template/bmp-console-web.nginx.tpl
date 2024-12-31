log_format  main1  '$time_local | $remote_addr | $upstream_addr | $http_host | $status | $body_bytes_sent | $upstream_response_time | $request_time | $http_host | $request | $http_referer';

server {
    listen 8080;
    server_name bmp-console.bmp.local;
    charset utf-8;
    access_log  /var/log/nginx/access.log  main1;

    error_page 500 502 503 504 /50x.html;
    location = /50x.html {
        root html;
    }

    root /home/bmp/bmp-console-web;
    index index.html;

    location /console-web/ {
        resolver 127.0.0.11 ipv6=off valid=2s; 
        set $bmp_console_api "bmp-console-api:8800";
        rewrite ^/console-web/(.*) /$1 break;
        proxy_pass http://$bmp_console_api;
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