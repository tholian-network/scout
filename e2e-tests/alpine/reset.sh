#!/bin/sh
set -e

apk add --allow-untrusted --force-broken-world /var/cache/e2e/apache2-old.apk >/tmp/e2e-reset.log 2>&1 || {
    cat /tmp/e2e-reset.log
    exit 1
}
