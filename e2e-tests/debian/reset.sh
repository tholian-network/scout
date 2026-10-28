#!/bin/sh
set -e

dpkg -i /var/cache/e2e/*.deb >/tmp/e2e-reset.log 2>&1 || {
    cat /tmp/e2e-reset.log
    exit 1
}
