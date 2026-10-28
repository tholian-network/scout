#!/bin/sh
set -e

/usr/sbin/httpd -DFOREGROUND >/tmp/e2e-service.log 2>&1 &
sleep 2

exec "$@"
