package color

import "runtime"

var Reset = "\033[0m"
var Error = "\033[31m"
var Success = "\033[32m"
var Warning = "\033[33m"
var Security = "\033[34m"
var Debug = "\033[37m"

// color reserve
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Error = ""
		Success = ""
		Warning = ""
		Security = ""
		Debug = ""
		White = ""
	}
}

// Colorize wraps a given message in a given color
func Colorize(color, message string) string {
	return color + message + Reset
}
