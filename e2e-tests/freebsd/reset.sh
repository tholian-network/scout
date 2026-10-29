#!/bin/sh
set -eu

export ASSUME_ALWAYS_YES=yes
export LC_ALL=C

PKG=/usr/local/sbin/pkg
if [ ! -x "$PKG" ]; then
    PKG=pkg
fi

$PKG delete -y e2e-scout-pkg || true
$PKG add -f /var/cache/e2e/repo/e2e-scout-pkg-1.0.pkg
$PKG update -f
