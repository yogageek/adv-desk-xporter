package util

import (
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

func OpenLog() bool {
	EnvlogInfo := os.Getenv("LOGINFO")
	return EnvlogInfo == "true"
}

func PrintGreen(s interface{}) {
	if !OpenLog() {
		return
	}
	fmt.Println(Green("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Green(ps))
}

func PrintBlue(s interface{}) {
	if !OpenLog() {
		return
	}
	fmt.Println(Blue("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Blue(ps))
}

func PrintYellow(s interface{}) {
	if !OpenLog() {
		return
	}
	fmt.Println(Yellow("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Yellow(ps))
}

func PrintCyan(s interface{}) {
	if !OpenLog() {
		return
	}
	fmt.Println(Cyan("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Cyan(ps))
}
