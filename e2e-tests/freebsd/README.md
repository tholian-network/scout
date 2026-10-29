# FreeBSD End-to-End Tests

This target runs the full Scout inventory suite inside a pinned FreeBSD
15.1-RELEASE QEMU virtual machine.

Podman cannot run FreeBSD containers on a Linux host because containers share
the host kernel. The FreeBSD target therefore uses QEMU full-system emulation
instead of the Podman test matrix.

## Prerequisites

- `qemu-system-x86_64`
- `qemu-img`
- `xz`
- `ssh`, `scp`, and `ssh-keygen`
- `genisoimage`, `mkisofs`, or `xorriso`
- KVM is preferred, but QEMU falls back to TCG when KVM is unavailable

## Image

The runner downloads and verifies:

```text
https://download.freebsd.org/releases/VM-IMAGES/15.1-RELEASE/amd64/Latest/FreeBSD-15.1-RELEASE-amd64-BASIC-CLOUDINIT-ufs.qcow2.xz
```

SHA256:

```text
e4ca4db889f8559c9b9dfcacc70405c038476f4b6d41649b152d3809a2ed9e1f
```

The cached image and run directories live under `build/tests/freebsd/`.

## Running

```bash
make test-freebsd
```

The runner cross-builds the FreeBSD test binary and the helper service, boots
the VM with a cloud-init seed ISO, configures a local two-version `pkg`
repository, installs `e2e-scout-pkg` 1.0, starts the helper service, and runs
the suite as root.

The VM is ephemeral. A fresh qcow2 overlay is created for each run, and the
serial log is kept under `build/tests/freebsd/runs/`.
