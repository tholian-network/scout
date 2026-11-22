package pip

import "os"

func isDirectory(path string) bool {

	stat, err := os.Stat(path)

	return err == nil && stat.IsDir()

}

func isFile(path string) bool {

	stat, err := os.Stat(path)

	return err == nil && stat.IsDir() == false

}
