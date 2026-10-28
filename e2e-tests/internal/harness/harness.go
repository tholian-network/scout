package harness

import "tholian-core/structs"
import "tholian-scout/actions"
import "os"
import "os/exec"
import "testing"

type Config struct {
	Package        string
	Service        string
	OldVersion     string
	PackageManager string
	UpdateManager  string
	Reset          string
}

func FromEnv(t *testing.T) Config {

	t.Helper()

	return Config{
		Package:        env(t, "E2E_PACKAGE"),
		Service:        env(t, "E2E_SERVICE"),
		OldVersion:     env(t, "E2E_OLD_VERSION"),
		PackageManager: env(t, "E2E_PACKAGE_MANAGER"),
		UpdateManager:  env(t, "E2E_UPDATE_MANAGER"),
		Reset:          env(t, "E2E_RESET"),
	}

}

func env(t *testing.T, key string) string {

	t.Helper()

	value := os.Getenv(key)

	if value == "" {
		t.Fatalf("Missing required environment variable %s", key)
	}

	return value

}

func NewConsole(t *testing.T) *structs.Console {

	t.Helper()

	file, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)

	if err != nil {
		t.Fatalf("Cannot open %s: %s", os.DevNull, err.Error())
	}

	t.Cleanup(func() {
		file.Close()
	})

	return structs.NewConsole(file, file, 0)

}

func NewSystem(t *testing.T) (*structs.Console, *structs.System) {

	t.Helper()

	console := NewConsole(t)
	system := actions.Init(console)

	if system == nil {
		t.Fatal("Expected Init() to return a System")
	}

	return console, system

}

func Reset(t *testing.T, config Config) {

	t.Helper()

	cmd := exec.Command(config.Reset)
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Cannot reset %s: %s\n%s", config.Package, err.Error(), string(output))
	}

}

func FindPackage(packages []structs.Package, name string) (structs.Package, bool) {

	for p := 0; p < len(packages); p++ {
		if packages[p].Name == name {
			return packages[p], true
		}
	}

	return structs.Package{}, false

}

func FindUpdate(updates []structs.Update, name string) (structs.Update, bool) {

	for u := 0; u < len(updates); u++ {
		if updates[u].Name == name {
			return updates[u], true
		}
	}

	return structs.Update{}, false

}

func FindAntique(antiques []structs.Antique, name string) (structs.Antique, bool) {

	for a := 0; a < len(antiques); a++ {
		if antiques[a].Name == name {
			return antiques[a], true
		}
	}

	return structs.Antique{}, false

}

func FindService(services []structs.Program, name string) (structs.Program, bool) {

	for s := 0; s < len(services); s++ {
		if services[s].Name == name {
			return services[s], true
		}
	}

	return structs.Program{}, false

}

func FindProgram(programs []structs.Program, name string) (structs.Program, bool) {

	for p := 0; p < len(programs); p++ {
		if programs[p].Name == name {
			return programs[p], true
		}
	}

	return structs.Program{}, false

}
