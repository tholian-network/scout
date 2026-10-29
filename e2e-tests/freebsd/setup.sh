#!/bin/sh
set -eu

export ASSUME_ALWAYS_YES=yes
export LC_ALL=C

mkdir -p /var/cache/e2e/repo /var/cache/e2e/stage/usr/local/bin /usr/local/etc/pkg/repos

if [ ! -x /usr/local/sbin/pkg ] && [ -x /usr/sbin/pkg ]; then
    /usr/sbin/pkg bootstrap -y || true
fi

PKG=/usr/local/sbin/pkg
if [ ! -x "$PKG" ]; then
    PKG=pkg
fi

for repo in FreeBSD FreeBSD-base FreeBSD-ports FreeBSD-ports-kmods; do
    printf '%s: { enabled: NO }\n' "$repo" > "/usr/local/etc/pkg/repos/$repo.conf"
done

cat > /usr/local/etc/pkg/repos/e2e.conf <<EOF
e2e: {
    url: "file:///var/cache/e2e/repo",
    enabled: true,
    signature_type: "none"
}
EOF

cp /tmp/e2e-scout-pkg /var/cache/e2e/stage/usr/local/bin/e2e-scout-pkg
chmod 755 /var/cache/e2e/stage/usr/local/bin/e2e-scout-pkg

ABI=$($PKG config abi 2>/dev/null || echo "FreeBSD:15:amd64")

write_manifest() {
    version=$1
    file=$2
    cat > "$file" <<EOF
name: "e2e-scout-pkg"
version: "$version"
origin: "sysutils/e2e-scout-pkg"
comment: "Scout FreeBSD E2E package"
desc: "Scout FreeBSD E2E package"
www: "https://example.com/e2e-scout-pkg"
maintainer: "e2e@example.com"
prefix: "/usr/local"
arch: "$ABI"
categories: [ "sysutils" ]
EOF
}

rm -rf /var/cache/e2e/repo
mkdir -p /var/cache/e2e/repo

write_manifest 1.0 /var/cache/e2e/manifest1
write_manifest 2.0 /var/cache/e2e/manifest2

printf '/usr/local/bin/e2e-scout-pkg\n' > /var/cache/e2e/plist

$PKG create -M /var/cache/e2e/manifest1 -p /var/cache/e2e/plist -r /var/cache/e2e/stage -o /var/cache/e2e/repo
$PKG create -M /var/cache/e2e/manifest2 -p /var/cache/e2e/plist -r /var/cache/e2e/stage -o /var/cache/e2e/repo
$PKG repo /var/cache/e2e/repo
$PKG update -f
$PKG add -f /var/cache/e2e/repo/e2e-scout-pkg-1.0.pkg

cp /tmp/e2e-reset.sh /var/cache/e2e/reset.sh
chmod 755 /var/cache/e2e/reset.sh

cd /
/usr/local/bin/e2e-scout-pkg --listen :8080 >/tmp/e2e-scout-pkg.log 2>&1 &
sleep 1

mkdir -p /var/log
printf '%s\n' "$(date '+%b %e %H:%M:%S') freebsd sshd[1234]: Failed password for invalid user test from 203.0.113.7 port 5555 ssh2" >> /var/log/auth.log

cp /bin/sleep /tmp/e2e-sleep
chmod 755 /tmp/e2e-sleep
/tmp/e2e-sleep 600 >/tmp/e2e-sleep.log 2>&1 &
