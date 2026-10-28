#!/bin/sh
set -e

pacman -U --noconfirm /var/cache/e2e/apache-old.pkg.tar.zst >/tmp/e2e-reset.log 2>&1 || {
    cat /tmp/e2e-reset.log
    exit 1
}
