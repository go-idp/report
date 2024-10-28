package os

import (
	"github.com/go-zoox/ip"
	"github.com/go-zoox/once"
)

func IP() string {
	return once.Get("ip", func() string {
		if v, err := ip.GetPublicIP(); err == nil {
			return v
		}

		return "unknown"
	})
}

func InternalIP() string {
	return once.Get("internal_ip", func() string {
		if v, err := ip.GetInternalIP(); err == nil {
			return v
		}

		return "unknown"
	})
}
