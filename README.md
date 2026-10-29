
# Tholian® Scout Inventory Agent

Tholian® Scout is the Inventory Agent for the Tholian Network. It collect system
reports including all necessary data and writes them to an external file.

## Building

```bash
# Build Dependencies
sudo pacman -S --needed go make;

cd /path/to/tholian-network/scout;

make check;
make build;
```

## Testing

```bash
# Test Dependencies
sudo pacman -S --needed go make podman;

cd /path/to/tholian-network/scout;

# Formatting, vet, and unit tests
make check;

# Containerized end-to-end tests (requires Podman)
make test;
```

`make test` builds `build/tests/<distribution>-<arch>`, builds the containers
from `./e2e-tests/<distribution>/Containerfile`, and runs the tests inside them.

## Usage

```bash
cd /path/to/tholian-network/scout;

cd ./build/linux;

sudo env THOLIAN_REPORT="/tmp/tholian-report.json" ./tholian-scout-amd64 collect;
```

## License

The Tholian Scout is licensed under the
[GNU Affero General Public License v3.0](LICENSE.txt).

`(c) 2020-2026 by Tholian(r) Network`

