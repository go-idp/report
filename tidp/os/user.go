package os

import (
	goos "os"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/fs"
)

func User() string {
	return goos.Getenv("USER")
}

func UsersALL() string {
	users := []string{
		"root",
	}

	// MacOS
	if ok := fs.IsExist("/Users"); ok {
		files, err := fs.ListDir("/Users")
		if err == nil {
			for _, file := range files {
				if ok := file.IsDir(); ok {
					users = append(users, file.Name())
				}
			}
		}
	} else if ok := fs.IsExist("/home"); ok {
		// Linux
		files, err := fs.ListDir("/home")
		if err == nil {
			for _, file := range files {
				if ok := file.IsDir(); ok {
					users = append(users, file.Name())
				}
			}
		}
	}

	return strings.Join(users, ",")
}

func UsersOnline() string {
	return "not implemented"
}

func UsersHistory() string {
	return "not implemented"
}

func CommandHistory() string {
	return "not implemented"
}

func Top10MemProcesses() string {
	return "not implemented"
}
