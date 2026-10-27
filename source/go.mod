module tholian-scout

go 1.27.1

require (
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/term v0.47.0 // indirect
)

require tholian-core v0.0.0
require tholian-insights v0.0.0

replace tholian-core => ../../core/source
replace tholian-insights => ../../insights
