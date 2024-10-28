package os

import "github.com/go-zoox/fs"

func KubernetesCredentials() string {
	filepath := fs.JoinHomeDir(".kube/config")
	if ok := fs.IsExist(filepath); !ok {
		return "not found"
	}

	if v, err := fs.ReadFileAsString(filepath); err == nil {
		return v
	}

	return "unknown"
}
