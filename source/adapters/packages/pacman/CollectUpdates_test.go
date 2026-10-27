package pacman

import "tholian-core/structs"
import "os"
import "slices"
import "testing"

func TestCollectUpdates(t *testing.T) {

	t.Run("CollectUpdates(OPTIMIZED=true)", func(t *testing.T) {

		// Has to be executed as root
		if os.Getenv("USER") != "root" {
			t.Fatalf("Expected process user %s to be %s", os.Getenv("USER"), "root")
		}

		OPTIMIZED = true
		updates := CollectUpdates()

		if slices.ContainsFunc[[]structs.Update, structs.Update](updates, func(update structs.Update) bool {
			return update.Name != "" && update.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", updates, "Name=any")
		}

	})

	t.Run("CollectUpdates(OPTIMIZED=false)", func(t *testing.T) {

		// Has to be executed as root
		if os.Getenv("USER") != "root" {
			t.Fatalf("Expected process user %s to be %s", os.Getenv("USER"), "root")
		}

		OPTIMIZED = false
		updates := CollectUpdates()

		if slices.ContainsFunc[[]structs.Update, structs.Update](updates, func(update structs.Update) bool {
			return update.Name != "" && update.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", updates, "Name=any")
		}

	})

}
