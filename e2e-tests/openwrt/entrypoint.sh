#!/bin/sh
set -e

cp /bin/busybox /tmp/busybox
chmod +x /tmp/busybox
/tmp/busybox sleep 300 &

exec "$@"
