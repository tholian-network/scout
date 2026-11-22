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

	AssembleProgramPackages(&program)

	if len(program.Packages) != 3 {
		t.Fatalf("Expected 3 program packages, got %d", len(program.Packages))
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 3 {
		t.Fatalf("Expected packages to be deduplicated, got %d", len(program.Packages))
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
