package apt

import "testing"

func TestCollectUpdate(t *testing.T) {

	t.Run("CollectUpdate(coreutils)", func(t *testing.T) {

		coreutils := CollectUpdate("coreutils", "*")

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

	})

	t.Run("CollectUpdate(libc6)", func(t *testing.T) {

		libc := CollectUpdate("libc6", "*")

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

	})

}
