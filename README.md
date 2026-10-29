
# Tholian® Scout Inventory Agent

Tholian® Scout is the Inventory Agent for the Tholian Network. It collect system
reports including all necessary data and writes them to an external file.

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

