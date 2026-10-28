module tholian-scout-e2e

go 1.27.1

require (
	tholian-core v0.0.0
	tholian-scout v0.0.0
)

require (
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/term v0.47.0 // indirect
	tholian-insights v0.0.0 // indirect
)

replace tholian-core => ../../core/source

replace tholian-insights => ../../insights

replace tholian-scout => ../source
