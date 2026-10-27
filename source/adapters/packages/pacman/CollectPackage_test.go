package pacman

import "tholian-core/matchers"
import "slices"
import "testing"

func TestCollectPackage(t *testing.T) {

	t.Run("CollectPackage(coreutils)", func(t *testing.T) {

		coreutils := CollectPackage("coreutils")

		if coreutils.Name != "coreutils" {
			t.Errorf("Expected %s to be %s", coreutils.Name, "coreutils")
		}

		if coreutils.Version.IsValid() == false {
			t.Errorf("Expected %s to be valid", coreutils.Version.String())
		}

		if coreutils.Architecture.String() == "" {
			t.Errorf("Expected %s to be %s", coreutils.Architecture.String(), "x86_64")
		}

		if coreutils.Manager.String() != "pacman" {
			t.Errorf("Expected %s to be %s", coreutils.Manager.String(), "pacman")
		}

		if slices.Contains(coreutils.Filesystem, "/usr/lib/coreutils/libstdbuf.so") == false {
			t.Errorf("Expected %v to contain %s", coreutils.Filesystem, "/usr/lib/coreutils/libstdbuf.so")
		}

		if slices.ContainsFunc[[]matchers.Package, matchers.Package](coreutils.Dependencies, func(dependency matchers.Package) bool {
			return dependency.Name == "glibc"
		}) == false {
			t.Errorf("Expected %v to contain %s", coreutils.Dependencies, "Name=glibc")
		}

	})

	t.Run("CollectPackage(glibc)", func(t *testing.T) {

		glibc := CollectPackage("glibc")

		if glibc.Name != "glibc" {
			t.Errorf("Expected %s to be %s", glibc.Name, "glibc")
		}

		if glibc.Version.IsValid() == false {
			t.Errorf("Expected %s to be valid", glibc.Version.String())
		}

		if glibc.Architecture.String() == "" {
			t.Errorf("Expected %s to be %s", glibc.Architecture.String(), "x86_64")
		}

		if glibc.Manager.String() != "pacman" {
			t.Errorf("Expected %s to be %s", glibc.Manager.String(), "pacman")
		}

		if slices.Contains(glibc.Filesystem, "/usr/lib/libc.so") == false {
			t.Errorf("Expected %v to contain %s", glibc.Filesystem, "/usr/lib/libc.so")
		}

		if slices.ContainsFunc[[]matchers.Package, matchers.Package](glibc.Dependencies, func(dependency matchers.Package) bool {
			return dependency.Name == "filesystem"
		}) == false {
			t.Errorf("Expected %v to contain %s", glibc.Dependencies, "Name=filesystem")
		}

	})

}
