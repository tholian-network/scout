package fs

import "io"
import "os"

func CopySelf(path string) bool {

	source, err1 := os.OpenFile(os.Args[0], os.O_RDONLY, 0755)
	target, err2 := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err1 == nil && err2 == nil {

		_, err3 := io.Copy(target, source)

		if err3 == nil {

			source.Close()
			target.Close()

			return true

		}

	}

	return false

}
