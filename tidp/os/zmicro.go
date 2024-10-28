package os

import (
	"fmt"
	"strings"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/once"
)

func ZmicroVersion() string {
	return once.Get("zmicro_version", func() string {
		filepath := "/usr/local/lib/zmicro/mod"
		if ok := fs.IsExist(filepath); !ok {
			return "not found"
		}

		if v, err := fs.ReadFileLines(filepath); err == nil {
			if len(v) == 0 {
				return "unknown"
			}

			if len(v[0]) < 8 {
				return strings.TrimSpace(v[0])
			}

			return strings.TrimSpace(v[0][8:])
		}

		return "unknown"
	})
}

func ZmicroVersionDetail() string {
	return once.Get("zmicro_version_detail", func() string {
		version := ZmicroVersion()
		git_commit_hash := GitCommitHash("/usr/local/lib/zmicro")
		git_commit_timestamp := GitCommitTimestamp("/usr/local/lib/zmicro")

		return fmt.Sprintf("%s (%s %s)", version, git_commit_hash, git_commit_timestamp)
	})
}
