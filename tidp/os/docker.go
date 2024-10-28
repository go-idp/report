package os

import "github.com/go-zoox/fs"

func DockerCredentials() string {
	filepath := fs.JoinHomeDir(".docker/config.json")
	if ok := fs.IsExist(filepath); !ok {
		return "not found"
	}

	if v, err := fs.ReadFileAsString(filepath); err == nil {
		return v
	}

	return "unknown"
}
