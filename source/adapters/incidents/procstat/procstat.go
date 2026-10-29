package procstat

import programs_procstat "tholian-scout/adapters/programs/procstat"

var SUPPORTED bool = false

func init() {

	if programs_procstat.SUPPORTED == true {
		SUPPORTED = true
	}

}
