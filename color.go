package color

import "runtime"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	// ANSI Escape Codes funktionieren unter Windows nicht. Deshalb werden die Escape Codes hier neutralisiert.
	if runtime.GOOS == "windows" {
		Reset = "\033[0m"
		Red = "\033[31m"
		Green = "\033[32m"
		Yellow = "\033[33m"
		Blue = "\033[34m"
		Purple = "\033[35m"
		Cyan = "\033[36m"
		Gray = "\033[37m"
		White = "\033[97m"
	}
}

// Ize "wrapped" die uebergebene Nachricht in die angegebene Farbe
func Ize(color, message string) string {
	// Der Escape Code wird vorangesetzt und faerbt die Consolen-Zeilen ein, dann wird die Zeile ausgegeben und anschließen der Reset-Escape-Code gesendet, um die Kolorierung wieder zurueckzusetzen.
	return color + message + Reset
}
