package os

func IsMacOS() bool {
	return Kernel() == "Darwin"
}
