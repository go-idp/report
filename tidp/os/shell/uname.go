package shell

func Uname(args ...string) (result string, err error) {
	return Exec("uname", args...)
}
