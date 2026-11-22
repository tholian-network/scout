package npm

import "testing"
import "tholian-core/structs"

func TestAssembleProgram(t *testing.T) {

	SUPPORTED = true

	program := structs.NewProgram(1337, "node")
	program.SetCommand("/usr/bin/node")
	program.SetFolder(fixture(t, "project"))

	if IsProgram(program) == false {
		t.Fatal("Expected testdata/project to be recognized as npm program")
	}

	AssembleProgramDependencies(&program)

	if len(program.Dependencies) != 3 {
		t.Fatalf("Expected 3 program dependencies, got %d", len(program.Dependencies))
	}

	if program.Dependencies[0].Name != "@scope/pkg" ||
		program.Dependencies[1].Name != "left-pad" ||
		program.Dependencies[2].Name != "typescript" {
		t.Fatalf("Expected sorted program dependencies, got %v", program.Dependencies)
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 2 {
		t.Fatalf("Expected 2 program packages, got %d", len(program.Packages))
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 2 {
		t.Fatalf("Expected packages to be deduplicated, got %d", len(program.Packages))
	}

}

func TestAssembleProgramNegative(t *testing.T) {

	program := structs.NewProgram(1337, "node")
	program.SetCommand("/usr/bin/node")
	program.SetFolder(fixture(t, "invalid"))

	if IsProgram(program) == true {
		t.Fatal("Expected folder without node_modules not to be an npm program")
	}

}
