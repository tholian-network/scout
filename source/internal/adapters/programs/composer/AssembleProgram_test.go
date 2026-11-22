package composer

import "testing"
import "tholian-core/structs"

func TestAssembleProgram(t *testing.T) {

	SUPPORTED = true

	program := structs.NewProgram(1337, "php")
	program.SetCommand("/usr/bin/php")
	program.SetFolder(fixture(t, "project"))

	if IsProgram(program) == false {
		t.Fatal("Expected composer project to be recognized as program")
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 2 {
		t.Fatalf("Expected 2 program packages, got %d", len(program.Packages))
	}

	AssembleProgramPackages(&program)

	if len(program.Packages) != 2 {
		t.Fatalf("Expected packages to be deduplicated, got %d", len(program.Packages))
	}

	AssembleProgramDependencies(&program)

	if len(program.Dependencies) != 3 {
		t.Fatalf("Expected 3 program dependencies, got %d", len(program.Dependencies))
	}

	if program.Dependencies[0].Name != "monolog/monolog" ||
		program.Dependencies[1].Name != "symfony/console" ||
		program.Dependencies[2].Name != "phpunit/phpunit" {
		t.Fatalf("Expected deterministic program dependencies, got %v", program.Dependencies)
	}

	if program.Dependencies[0].Version != ">= 3.5" ||
		program.Dependencies[1].Version != ">= 6.4" ||
		program.Dependencies[2].Version != ">= 10.5" {
		t.Fatalf("Expected normalized program dependency versions, got %v", program.Dependencies)
	}

	if len(program.Unresolved) != 0 {
		t.Fatalf("Expected no unresolved program dependencies, got %d", len(program.Unresolved))
	}

}

func TestAssembleProgramNegative(t *testing.T) {

	program := structs.NewProgram(1337, "php")
	program.SetCommand("/usr/bin/php")
	program.SetFolder(fixture(t, "lock"))

	if IsProgram(program) == true {
		t.Fatal("Expected composer project without vendor directory not to be a program")
	}

}
