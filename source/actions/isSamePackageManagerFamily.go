package actions

func isSamePackageManagerFamily(first string, second string) bool {

	if first == second {
		return true
	}

	rpm_family := map[string]bool{
		"dnf":    true,
		"rpm":    true,
		"tdnf":   true,
		"zypper": true,
	}

	return rpm_family[first] == true && rpm_family[second] == true

}
