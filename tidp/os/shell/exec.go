package shell

import (
	"os/exec"
	"strings"
)

func Exec(command string, args ...string) (result string, err error) {
	cmd := exec.Command(command, args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
