package syslog

import "io"
import "os"

const MaxLogBytes = 16 * 1024 * 1024

func readSyslogFile(name string) ([]byte, error) {

	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return nil, err
	}

	size := stat.Size()
	start := int64(0)

	if size > MaxLogBytes {
		start = size - MaxLogBytes
	}

	if start > 0 {
		if _, err := file.Seek(start, io.SeekStart); err != nil {
			return nil, err
		}
	}

	buffer, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	if start > 0 {

		for b := 0; b < len(buffer); b++ {
			if buffer[b] == '\n' {
				buffer = buffer[b+1:]
				break
			}
		}

	}

	return buffer, nil

}
