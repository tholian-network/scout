package proc

import "tholian-core/caches"

var Programs *caches.Programs

func init() {
	Programs = caches.NewPrograms()
}
