#!/bin/sh
set -e

mkdir -p /run/httpd /var/run/httpd

/usr/bin/httpd -DFOREGROUND >/tmp/e2e-service.log 2>&1 &
sleep 2

exec "$@"
