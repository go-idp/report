package os

import (
	"github.com/go-idp/report/tidp/os/shell"
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/fs"
	"github.com/go-zoox/once"
)

func DeviceID() string {
	return once.Get("device_id", func() string {
		if IsMacOS() {
			output, err := shell.Exec("sh", "-c", `ioreg -rd1 -c IOPlatformExpertDevice | grep "IOPlatformUUID" | awk -F '"' '{print $4}'`)
			if err == nil {
				return strings.TrimSpace(output)
			}

			output, err = shell.Exec("sh", "-c", `system_profiler SPHardwareDataType | grep "Hardware UUID" | awk -F ':' '{print $2}'`)
			if err == nil {
				return strings.TrimSpace(output)
			}
		}

		if fs.IsExist("/var/lib/dbus/machine-id") {
			if v, err := fs.ReadFileAsString("/var/lib/dbus/machine-id"); err == nil {
				return v
			}
		}

		if fs.IsExist("/etc/machine-id") {
			if v, err := fs.ReadFileAsString("/etc/machine-id"); err == nil {
				return v
			}
		}

		if fs.IsExist(fs.HomeDir() + "/.machine-id.zmicro") {
			if v, err := fs.ReadFileAsString(fs.HomeDir() + "/.machine-id.zmicro"); err == nil {
				return v
			}
		}

		if fs.IsExist("/proc/sys/kernel/random/boot_id") {
			if v, err := fs.ReadFileAsString("/proc/sys/kernel/random/boot_id"); err == nil {
				return v
			}
		}

		return ""
	})
}
