package pacman

import "testing"

func TestCollectUpdate(t *testing.T) {

	t.Run("CollectUpdate(coreutils)", func(t *testing.T) {

		coreutils := CollectUpdate("coreutils")

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

	})

	t.Run("CollectUpdate(glibc)", func(t *testing.T) {

		glibc := CollectUpdate("glibc")

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

	})

}
