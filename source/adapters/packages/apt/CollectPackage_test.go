package apt

import "tholian-core/matchers"
import "slices"
import "testing"

func TestCollectPackage(t *testing.T) {

	t.Run("CollectPackage(coreutils)", func(t *testing.T) {

		coreutils := CollectPackage("coreutils", "*")

		if coreutils.Name != "coreutils" {
			t.Errorf("Expected %s to be %s", coreutils.Name, "coreutils")
		}

		if coreutils.Version.IsValid() == false {
			t.Errorf("Expected %s to be valid", coreutils.Version.String())
		}

		if coreutils.Architecture.String() == "" {
			t.Errorf("Expected %s to be %s", coreutils.Architecture.String(), "x86_64")
		}

		if coreutils.Manager.String() != "apt" {
			t.Errorf("Expected %s to be %s", coreutils.Manager.String(), "apt")
		}

		if slices.Contains(coreutils.Filesystem, "/usr/libexec/coreutils/libstdbuf.so") == false {
			t.Errorf("Expected %v to contain %s", coreutils.Filesystem, "/usr/libexec/coreutils/libstdbuf.so")
		}

		if slices.ContainsFunc[[]matchers.Package, matchers.Package](coreutils.Dependencies, func(dependency matchers.Package) bool {
			return dependency.Name == "libc6"
		}) == false {
			t.Errorf("Expected %v to contain %s", coreutils.Dependencies, "Name=libc6")
		}

	})

	t.Run("CollectPackage(libc6)", func(t *testing.T) {

		libc := CollectPackage("libc6", "*")

		if libc.Name != "libc6" {
			t.Errorf("Expected %s to be %s", libc.Name, "libc6")
		}

		if libc.Version.IsValid() == false {
			t.Errorf("Expected %s to be valid", libc.Version.String())
		}

		if libc.Architecture.String() == "" {
			t.Errorf("Expected %s to be %s", libc.Architecture.String(), "x86_64")
		}

		if libc.Manager.String() != "apt" {
			t.Errorf("Expected %s to be %s", libc.Manager.String(), "apt")
		}

		if slices.Contains(libc.Filesystem, "/usr/lib/x86_64-linux-gnu/libc.so") == false {
			t.Errorf("Expected %v to contain %s", libc.Filesystem, "/usr/lib/x86_64-linux-gnu/libc.so")
		}

		if slices.ContainsFunc[[]matchers.Package, matchers.Package](libc.Dependencies, func(dependency matchers.Package) bool {
			return dependency.Name == "libgcc-s1"
		}) == false {
			t.Errorf("Expected %v to contain %s", libc.Dependencies, "Name=libgcc-s1")
		}

	})

}
