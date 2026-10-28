#!/bin/sh
set -e

mkdir -p /var/lock /var/run

opkg install --force-downgrade --force-reinstall --force-overwrite /var/cache/e2e/feed/uhttpd_1.0_x86_64.ipk >/tmp/e2e-reset.log 2>&1 || {
    cat /tmp/e2e-reset.log
    exit 1
}
