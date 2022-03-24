#!/bin/bash
# setup nginx
mkdir -p /data/nginx/cache
mv /nginx_config /etc/nginx/sites-available/default
B64KEY=$(echo -n "$INPUT_RPROXYUSER:$INPUT_RPROXYKEY" | base64 -w0)
AUTHLINE=$(echo -n "Basic $B64KEY")
sed -i "s/REPLACE/$AUTHLINE/g" /etc/nginx/sites-available/default
# service nginx restart
/usr/sbin/nginx -g "daemon off;"
