package actions

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "strconv"

func SearchUsers(console *structs.Console, system *structs.System, search matchers.User) []types.User {

	if system == nil {
		system = initSystem(console)
	}

	CollectUsers(console, system)

	console.Group("actions/SearchUsers")

	filtered := make([]types.User, 0)

	for u := 0; u < len(system.Users); u++ {

		user := system.Users[u]

		if search.Matches(user.Name, user.Password, user.Type) {
			filtered = append(filtered, user)
		}

	}

	console.Log("Searched " + strconv.Itoa(len(filtered)) + "/" + strconv.Itoa(len(system.Users)) + " Users")
	console.GroupEnd("actions/SearchUsers")

	return filtered

}
