#!/bin/sh
set -e

mkdir -p /var/lock /var/run /var/cache/e2e/feed
sed -i '/check_signature/d' /etc/opkg.conf

build_pkg() {
    version="$1"
    file="$2"
    dir="/tmp/opkg-$version"

    rm -rf "$dir"
    mkdir -p "$dir/control" "$dir/data/usr/sbin"

    cat > "$dir/control/control" <<EOF
Package: uhttpd
Version: $version
Architecture: x86_64
Maintainer: E2E <e2e@example.com>
Section: net
Priority: optional
Installed-Size: 1
Description: E2E uhttpd
EOF

    printf '#!/bin/sh\nexit 0\n' > "$dir/data/usr/sbin/uhttpd"
    chmod +x "$dir/data/usr/sbin/uhttpd"
    printf '2.0\n' > "$dir/debian-binary"

    tar -czf "$dir/control.tar.gz" -C "$dir/control" .
    tar -czf "$dir/data.tar.gz" -C "$dir/data" .
    tar -czf "$file" -C "$dir" ./debian-binary ./control.tar.gz ./data.tar.gz
}

build_pkg 1.0 /var/cache/e2e/feed/uhttpd_1.0_x86_64.ipk
build_pkg 2.0 /var/cache/e2e/feed/uhttpd_2.0_x86_64.ipk

cd /var/cache/e2e/feed
rm -f Packages Packages.gz

for version in 1.0 2.0; do
    file="uhttpd_${version}_x86_64.ipk"
    size=$(wc -c < "$file")
    sha=$(sha256sum "$file" | cut -d' ' -f1)

    cat >> Packages <<EOF
Package: uhttpd
Version: $version
Architecture: x86_64
Installed-Size: 1
Filename: $file
Size: $size
SHA256sum: $sha
Description: E2E uhttpd

EOF
done

gzip -kf Packages

printf 'src/gz e2e file:///var/cache/e2e/feed\n' > /etc/opkg/distfeeds.conf

opkg update
opkg install --force-downgrade --force-reinstall --force-overwrite /var/cache/e2e/feed/uhttpd_1.0_x86_64.ipk
