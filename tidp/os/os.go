package os

import (
	"os"
	"runtime"

	"github.com/go-idp/report/tidp/os/shell"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/once"
)

func Shell() string {
	return os.Getenv("SHELL")
}

func Home() string {
	return fs.HomeDir()
}

func Hostname() string {
	return once.Get("hostname", func() string {
		hostname, _ := os.Hostname()
		return hostname
	})
}

func Kernel() string {
	return once.Get("kernel", func() string {
		if v, err := shell.Uname("-s"); err == nil {
			return v
		}

		return "unknown"
	})
}

func Distribution() string {
	return once.Get("distribution", func() string {
		if ok := fs.IsExist("/etc/fedora-release"); ok {
			if v, err := fs.ReadFileAsString("/etc/fedora-release"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/redhat-release"); ok {
			if v, err := fs.ReadFileAsString("/etc/redhat-release"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/SuSE-release"); ok {
			if v, err := fs.ReadFileAsString("/etc/SuSE-release"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/mandrake-release"); ok {
			if v, err := fs.ReadFileAsString("/etc/mandrake-release"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/debian_version"); ok {
			if v, err := fs.ReadFileAsString("/etc/debian_version"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/gentoo-release"); ok {
			if v, err := fs.ReadFileAsString("/etc/gentoo-release"); err == nil {
				return v
			}
		}

		if ok := fs.IsExist("/etc/issue"); ok {
			if v, err := fs.ReadFileAsString("/etc/issue"); err == nil {
				return v
			}
		}

		return "unknown"
	})
}

func Arch() string {
	return runtime.GOARCH
}
