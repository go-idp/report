package shell

func Ls(args ...string) (result string, err error) {
	return Exec("ls", args...)
}
