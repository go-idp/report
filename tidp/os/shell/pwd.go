package shell

func Pwd() (result string, err error) {
	return Exec("pwd")
}
