
# Tholian® Scout Inventory Agent

Tholian® Scout is the Inventory Agent for the Tholian Network. It collect system
reports including all necessary data and writes them to an external file.

## Building

```bash
# Build Dependencies
sudo pacman -S --needed go;

cd /path/to/tholian-network/scout;

make check;
make build;
```

## Testing

```bash
cd /path/to/tholian-network/scout;

make test;
```

## Usage

```bash
cd /path/to/tholian-network/scout;

cd ./build/linux;

sudo env THOLIAN_REPORT="/tmp/tholian-report.json" ./tholian-scout-amd64 collect;
```

## License

Proprietary License. All Rights Reserved.

`(c) 2020-2026 by Tholian(r) Network`

