#!/bin/sh
set -e

mkdir -p /run/apache2 /var/run/apache2

. /etc/apache2/envvars

/usr/sbin/apache2 -DFOREGROUND >/tmp/e2e-service.log 2>&1 &
sleep 2

exec "$@"
