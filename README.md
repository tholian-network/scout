
# Tholian® Scout Inventory Agent

Tholian® Scout is the Inventory Agent for the Tholian Network. It collect system
reports including all necessary data and writes them to an external file.

## Platform Support

Scout detects distributions through `/etc/os-release` and collects inventory
through Linux kernel interfaces. The following distributions and package
managers are supported:

| Distribution             | Package Manager | Antiques | Packages | Devices | Drives | Incidents | Networks | Programs | Services | Updates | Users |
|--------------------------|-----------------|----------|----------|---------|--------|-----------|----------|----------|----------|---------|-------|
| AlmaLinux                | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Alpine Linux             | `apk`           | yes      | yes      | yes     | yes    | partial   | yes      | yes      | yes      | yes     | yes   |
| Amazon Linux             | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Antergos                 | `pacman`        | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Arch Linux               | `pacman`        | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| CBL-Mariner              | `tdnf`          | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| CentOS                   | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Debian                   | `apt`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Fedora                   | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Linux Mint               | `apt`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Manjaro                  | `pacman`        | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| OpenWrt                  | `opkg`          | yes      | yes      | yes     | yes    | partial   | yes      | yes      | yes      | yes     | yes   |
| openSUSE                 | `zypper`        | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Oracle Linux             | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Photon OS                | `tdnf`          | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Red Hat Enterprise Linux | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Rocky Linux              | `dnf`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| SUSE Linux Enterprise    | `zypper`        | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Trisquel                 | `apt`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |
| Ubuntu                   | `apt`           | yes      | yes      | yes     | yes    | yes       | yes      | yes      | yes      | yes     | yes   |

`yes` means a collector is implemented for that datatype. `no` means it is not
implemented. `partial` means incidents come from live process analysis and from
syslog sources when available, but historical journal entries are not collected.

Incidents are collected through `systemd`/`journalctl` where available, through
`procfs` process analysis on Linux, and through syslog files or OpenWrt
`logread`. On Alpine Linux and OpenWrt, live process findings and syslog-based
authentication/reconnaissance incidents are reported when the relevant logs or
services exist.

### Fallback Support

- On Debian-family systems without `apt`, Scout falls back to `dpkg`.
- On RPM-family systems without `dnf`, `tdnf`, or `zypper`, Scout falls back to `rpm`.

Both fallbacks collect packages only. FreeBSD, NetBSD, NixOS, and Gentoo are
recognized by the core distribution database, but their `pkg`, `pkgsrc`, `nix`,
and Portage adapters are not implemented yet.

## Requirements

The module is built using Go's module system. In the Tholian Network, all
repositories are checked out next to each other.

The shared Tholian Core Type System lives in the [tholian-core](https://github.com/tholian-network/core)
repository. This repository expects it to be checked out next to the scout folder:

```text
/path/to/Software/tholian-network/
|- core/     # tholian-core module
\- scout/    # this repository
```

`source/go.mod` resolves the dependency with `replace tholian-core => ../../core/source`,
which is relative to the `source/` directory. If your checkout uses a different directory
name or location, adjust that `replace` directive accordingly.

## Building

Building requires `make` and `go`.

```bash
# Install dependencies
sudo pacman -S --needed go make;

cd /path/to/tholian-network/scout;

make check;
make build;
```

## Usage

```bash
cd /path/to/tholian-network/scout;

cd ./build/linux;
sudo env THOLIAN_REPORT="/tmp/tholian-report.json" ./tholian-scout-amd64 collect;
```

## Testing

Testing requires `make`, `go`, and `podman`.

```bash
# Install dependencies
sudo pacman -S --needed go make podman;

cd /path/to/tholian-network/scout;

# Formatting, vet, and unit tests
make check;

# Containerized end-to-end tests (requires Podman)
make test;
```

`make test` does the following:

1. builds `build/tests/<distribution>-<arch>`
2. builds the containers from `./e2e-tests/<distribution>/Containerfile`
3. runs the previously built test binaries inside their containers

## License

The Tholian Scout is licensed under the
[GNU Affero General Public License v3.0](LICENSE.txt).

`(c) 2020-2026 by Tholian(r) Network`

