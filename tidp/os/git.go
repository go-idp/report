package os

import (
	"github.com/go-idp/report/tidp/os/shell"
	"github.com/go-zoox/fs"
)

func GitCredentials() string {
	filepath := fs.JoinHomeDir(".git-credentials")
	if ok := fs.IsExist(filepath); !ok {
		return "not found"
	}

	if v, err := fs.ReadFileAsString(filepath); err == nil {
		return v
	}

	return "unknown"
}

func GitCommitHash(path ...string) string {
	pathX := ""
	if len(path) > 0 {
		pathX = path[0]
	} else {
		pathX = fs.CurrentDir()
	}

	if v, err := shell.Exec("git", "-C", pathX, "rev-parse", "--short=9", "HEAD"); err == nil {
		return v
	}

	return ""
}

func GitCommitTimestamp(path ...string) string {
	pathX := ""
	if len(path) > 0 {
		pathX = path[0]
	} else {
		pathX = fs.CurrentDir()
	}

	if v, err := shell.Exec("git", "-C", pathX, "show", "-s", "--format=%cd", `--date="format:%Y-%m-%d_%H:%M:%S"`); err == nil {
		return v
	}

	return ""
}
