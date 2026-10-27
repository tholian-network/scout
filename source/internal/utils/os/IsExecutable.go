package os

import "os"
import "os/user"
import "strconv"
import "syscall"

func IsExecutable(file string) bool {

	stat, err1 := os.Stat(file)
	user, err2 := user.Current()

	if err1 == nil && err2 == nil && stat.IsDir() == false {

		sys_stat := stat.Sys().(*syscall.Stat_t)
		owner_uid := strconv.Itoa(int(sys_stat.Uid))
		owner_gid := strconv.Itoa(int(sys_stat.Gid))
		owner_mod := stat.Mode().Perm()

		if owner_uid == user.Uid && owner_mod&0100 != 0 {

			return true

		} else if owner_uid != user.Uid && owner_mod&0010 != 0 {

			groups, err3 := user.GroupIds()

			var is_same_group bool = false

			if err3 == nil {

				for g := 0; g < len(groups); g++ {

					if owner_gid == groups[g] {
						is_same_group = true
						break
					}

				}

			}

			if is_same_group == true {
				return true
			}

		} else if owner_uid != user.Uid && owner_mod&0001 != 0 {

			groups, err3 := user.GroupIds()

			var is_same_group bool = false

			if err3 == nil {

				for g := 0; g < len(groups); g++ {

					if owner_gid == groups[g] {
						is_same_group = true
						break
					}

				}

			}

			if is_same_group == false {
				return true
			}

		}

	}

	return false

}
