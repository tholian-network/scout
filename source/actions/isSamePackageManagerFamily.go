package actions

func isSamePackageManagerFamily(first string, second string) bool {

	if first == second {
		return true
	}

	dpkg_family := map[string]bool{
		"apt":  true,
		"dpkg": true,
	}

	if dpkg_family[first] == true && dpkg_family[second] == true {
		return true
	}

	rpm_family := map[string]bool{
		"dnf":    true,
		"rpm":    true,
		"tdnf":   true,
		"zypper": true,
	}

	if rpm_family[first] == true && rpm_family[second] == true {
		return true
	}

	return false

}
