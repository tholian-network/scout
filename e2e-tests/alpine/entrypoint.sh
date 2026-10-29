#!/bin/sh
set -e

/usr/sbin/httpd -DFOREGROUND >/tmp/e2e-service.log 2>&1 &
sleep 2

mkdir -p /var/log
busybox syslogd -O /var/log/messages >/tmp/e2e-syslog.log 2>&1 || true
sleep 1
busybox logger -t sshd "Failed password for root from 203.0.113.7 port 22 ssh2"

cp /bin/busybox /tmp/busybox
chmod +x /tmp/busybox
/tmp/busybox sleep 300 &

exec "$@"
