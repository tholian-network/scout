package pip

import "testing"
import "tholian-core/structs"

func TestAssembleProgram(t *testing.T) {

	SUPPORTED = true

	program := structs.NewProgram(1337, "python3")
	program.SetCommand("/usr/bin/python3")
	program.SetFolder(fixture(t, "project"))

	if IsProgram(program) == false {
		t.Fatal("Expected Python project to be recognized as program")
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 3 {
		t.Fatalf("Expected 3 program packages, got %d", len(program.Packages))
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 3 {
		t.Fatalf("Expected packages to be deduplicated, got %d", len(program.Packages))
	}

	AssembleProgramDependencies(&program)

	if len(program.Dependencies) != 7 {
		t.Fatalf("Expected 7 program dependencies, got %d", len(program.Dependencies))
	}

	if program.Dependencies[0].Name != "Flask" ||
		program.Dependencies[1].Name != "click" ||
		program.Dependencies[2].Name != "django" ||
		program.Dependencies[3].Name != "editable-pkg" ||
		program.Dependencies[4].Name != "local-pkg" ||
		program.Dependencies[5].Name != "requests" ||
		program.Dependencies[6].Name != "uvicorn" {
		t.Fatalf("Expected sorted program dependencies, got %v", program.Dependencies)
	}

	if program.Dependencies[0].Version != ">= 3.0" ||
		program.Dependencies[1].Version != "= 8.1.7" ||
		program.Dependencies[2].Version != "= 5.0.2" ||
		program.Dependencies[3].Version != "any" ||
		program.Dependencies[4].Version != "any" ||
		program.Dependencies[5].Version != "= 2.31.0" ||
		program.Dependencies[6].Version != ">= 0.29.0" {
		t.Fatalf("Expected normalized program dependency versions, got %v", program.Dependencies)
	}

	if len(program.Unresolved) != 0 {
		t.Fatalf("Expected no unresolved program dependencies, got %d", len(program.Unresolved))
	}

}

func TestIsProgramVenv(t *testing.T) {

	program := structs.NewProgram(1337, "gunicorn")
	program.SetCommand("/usr/local/bin/gunicorn")
	program.SetFolder(fixture(t, "project"))

	if IsProgram(program) == false {
		t.Fatal("Expected venv to identify Python program")
	}

}

func TestAssembleProgramNegative(t *testing.T) {

	program := structs.NewProgram(1337, "sh")
	program.SetCommand("/usr/bin/sh")
	program.SetFolder(fixture(t, "hostile"))

	if IsProgram(program) == true {
		t.Fatal("Expected shell without Python evidence not to be a Python program")
	}

}
