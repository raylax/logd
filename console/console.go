package console

import (
	"github.com/raylax/logd/model"
)

const (
	ColorBlack  string = "\u001b[30m"
	ColorRed           = "\u001b[31m"
	ColorGreen         = "\u001b[32m"
	ColorYellow        = "\u001b[33m"
	ColorBlue          = "\u001b[34m"
	ColorReset         = "\u001b[0m"
)

func Black(str string) string {
	return ColorBlack + str + ColorReset
}

func Red(str string) string {
	return ColorRed + str + ColorReset
}

func Green(str string) string {
	return ColorGreen + str + ColorReset
}

func Yellow(str string) string {
	return ColorYellow + str + ColorReset
}

func Blue(str string) string {
	return ColorBlue + str + ColorReset
}

func Println(line *model.LogLine)  {
	println(Yellow(line.Server) + " | " + Blue(line.File) + " | " + line.Line)
}

