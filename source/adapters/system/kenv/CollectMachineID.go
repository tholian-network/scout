package kenv

func CollectMachineID() string {

	var result string

	environment := CollectKernelEnvironment()

	if value, ok := environment["smbios.system.uuid"]; ok == true {
		result = value
	}

	return result

}
