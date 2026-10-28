# End-to-End Testing

The end-to-end tests live in `./e2e-tests/<target>` and run inside Podman
containers. Each container installs a package, downgrades it to a pinned
version, starts the expected service, and then runs the test binary against the
host inventory APIs.

## Test Targets

| Target | Base image | Package manager | Pinned package | Pinned version | Service | Expected antique | Expected update | Reset mechanism |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `almalinux` | `almalinux:9` | `dnf` | `httpd-core` | `2.4.57-11.el9_4.1` | `httpd` | `httpd-core` via `httpd` | `> 2.4.57-11.el9_4.1` | `rpm -U --oldpackage` |
| `alpine` | `alpine:3.21` | `apk` | `apache2` | `2.4.67-r0` | `httpd` | `apache2` via `httpd` | `> 2.4.67-r0` | `apk add` old package |
| `archlinux` | `archlinux:latest` | `pacman` | `apache` | `2.4.62-1` | `httpd` | `apache` via `httpd` | `> 2.4.62-1` | `pacman -U` old package |
| `cblmariner` | `mcr.microsoft.com/cbl-mariner/base/core:2.0` | `tdnf` | `httpd` | `2.4.62-1.cm2` | `httpd` | `httpd` via `httpd` | `> 2.4.62-1.cm2` | `rpm -U --oldpackage` |
| `debian` | `debian:bookworm-slim` | `apt` | `apache2-bin` | `2.4.62-1~deb12u1` | `apache2` | `apache2-bin` via `apache2` | `> 2.4.62-1~deb12u1` | `dpkg -i` old packages |
| `fedora` | `fedora:43` | `dnf` | `httpd-core` | `2.4.63-4.fc43` | `httpd` | `httpd-core` via `httpd` | `> 2.4.63-4.fc43` | `rpm -U --oldpackage` |
| `opensuse` | `opensuse/leap:15.6` | `zypper` | `apache2-prefork` | `2.4.58-150600.3.1` | `httpd-prefork` | `apache2-prefork` via `httpd-prefork` | `> 2.4.58-150600.3.1` | `rpm -U --oldpackage` |
| `rockylinux` | `rockylinux:9` | `dnf` | `httpd-core` | `2.4.57-11.el9_4.1` | `httpd` | `httpd-core` via `httpd` | `> 2.4.57-11.el9_4.1` | `rpm -U --oldpackage` |
| `rpm` | `almalinux:9-minimal` | `rpm` | `httpd-core` | `2.4.57-11.el9_4.1` | not used | not tested | not tested | none |

`Expected update` means any version newer than the pinned old version that the
distribution repository currently provides. The exact new version is not
pinned because repositories move over time.

The `rpm` target is a package-collection-only test case. It runs only
`CollectPackages` and `SearchPackages`, because the rpm adapter does not
implement updates.

## Pinned Package Sources

| Target | Pinned artifacts | Source |
| --- | --- | --- |
| `almalinux` | `httpd`, `httpd-core`, `httpd-tools`, `mod_lua`, `httpd-filesystem` | AlmaLinux 9.4 vault AppStream |
| `alpine` | `apache2` | Alpine v3.20 main repository |
| `archlinux` | `apache` | Arch Linux package archive |
| `cblmariner` | `httpd` | CBL-Mariner 2.0 base repository |
| `debian` | `apache2`, `apache2-bin`, `apache2-data`, `apache2-utils` | Debian snapshot `20240817T023924Z` |
| `fedora` | `httpd`, `httpd-core`, `httpd-tools`, `mod_lua`, `httpd-filesystem` | Fedora Koji `httpd-2.4.63-4.fc43` |
| `opensuse` | `apache2`, `apache2-prefork`, `apache2-utils` | openSUSE Leap 15.6 OSS repository |
| `rockylinux` | `httpd`, `httpd-core`, `httpd-tools`, `mod_lua`, `httpd-filesystem` | Rocky Linux 9.4 vault AppStream |
| `rpm` | `httpd`, `httpd-core`, `httpd-tools`, `mod_lua`, `httpd-filesystem` | AlmaLinux 9.4 vault AppStream |

## Expected Test Behavior

Full test targets (`almalinux`, `alpine`, `archlinux`, `cblmariner`, `debian`,
`fedora`, `opensuse`, `rockylinux`) verify:

- `CollectPackages` reports the pinned package at the pinned old version using
  the expected package manager.
- `CollectUpdates` reports an update for the pinned package with the same
  package manager and a version newer than the pinned old version.
- `CollectAntiques` and `SearchAntiques` find an antique for the pinned package
  and expected service.
- `Update`, `UpdatePackage`, and `UpdatePackages` upgrade the package and leave
  no remaining update for it.
- The service tests find the expected listening service.
- The generic device, drive, incident, network, program, and user tests pass or
  skip when the container runtime does not expose the required resources.

The `rpm` target verifies only:

- `CollectPackages` reports `httpd-core` at the pinned old version using the
  `rpm` package manager.
- `SearchPackages` finds `httpd-core`.

## Running Tests

```bash
# Run every test target
make test

# Run one target
make test TEST_DISTRIBUTIONS="cblmariner"

# Run multiple targets
make test TEST_DISTRIBUTIONS="cblmariner rpm"

# Build test binaries without running containers
make build-tests TEST_DISTRIBUTIONS="cblmariner rpm"
```

## Maintaining This Document

When changing a test target's `Containerfile`, entrypoint, or reset logic:

1. Update the matching test directory under `./e2e-tests`.
2. Update the `E2E_*` environment values in its `Containerfile`.
3. Update the tables in this document.
4. Run `make test TEST_DISTRIBUTIONS="<target>"`.
