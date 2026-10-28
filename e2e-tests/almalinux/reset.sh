#!/bin/sh
set -e

rpm -U --oldpackage --replacepkgs /var/cache/e2e/*.rpm >/tmp/e2e-reset.log 2>&1 || {
    cat /tmp/e2e-reset.log
    exit 1
}
