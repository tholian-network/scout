package apt

import "tholian-core/structs"
import "slices"
import "testing"

func TestCollectUpdates(t *testing.T) {

	t.Run("CollectUpdates()", func(t *testing.T) {

		updates := CollectUpdates()

		if slices.ContainsFunc[[]structs.Update, structs.Update](updates, func(update structs.Update) bool {
			return update.Name != "" && update.Version.IsValid()
		}) == false {
			t.Errorf("Expected %v to contain %s", updates, "Name=any")
		}

	})

}
