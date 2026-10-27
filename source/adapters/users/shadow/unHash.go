package shadow

import "strings"
import "fmt"

func unHash(password string) string {

	var result string = ""

	if strings.HasPrefix(password, "$") {

		tmp := strings.Split(password, "$")

		if len(tmp) == 4 {

			algo := tmp[1]
			salt := tmp[2]
			hash := tmp[3]

			if algo == "" || algo == "_" {
				// DES
				// should not be used
			} else if algo == "1" || algo == "md5" {
				// MD5
				// should not be used
			} else if algo == "2" || algo == "2a" || algo == "2b" || algo == "2x" || algo == "2y" {
				// bcrypt/blowfish
				// various broken implementations
				// should not be used
			} else if algo == "3" {
				// MD4 / NTHASH
				// should not be used
			} else if algo == "4" {
				// should not be used
			} else if algo == "5" {
				// SHA-256
			} else if algo == "6" {
				// SHA-512
			}

			// TODO: Implement unHashing of passwords
			fmt.Println(algo, salt, hash)

		}

	}

	return result

}
