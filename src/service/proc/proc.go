package proc

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func FindProcessByInode(inode uint32) (*int, error) {

	procFiles, err := ioutil.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	pids := []int{}

	for _, f := range procFiles {
		pid, err := strconv.Atoi(f.Name())
		if err == nil {
			pids = append(pids, pid)
		}
	}

	socketStr := fmt.Sprintf("socket:[%d]", inode)
	for _, pid := range pids {

		path := fmt.Sprintf("/proc/%d/fd/", pid)
		if _, err := os.Stat(path); err == nil {

			fds, err := ioutil.ReadDir(path)
			if err != nil {
				return nil, err
			}

			for _, fd := range fds {

				path := fmt.Sprintf("/proc/%d/fd/%s", pid, fd.Name())
				if _, err := os.Stat(path); err == nil {
					link, err := os.Readlink(path)
					if err != nil {
						return nil, err
					}

					if socketStr == link {
						return &pid, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("inode not found: %d", inode)
}

func GetProcessCommand(pid int) (string, error) {

	reader, err := os.Open(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
