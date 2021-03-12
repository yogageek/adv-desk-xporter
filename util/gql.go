package util

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
)

func PrintGreen(s interface{}) {
	fmt.Println(Green("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Green(ps))
}

func PrintBlue(s interface{}) {
	fmt.Println(Blue("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Blue(ps))
}

func PrintYellow(s interface{}) {
	fmt.Println(Yellow("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Yellow(ps))
}

func PrintCyan(s interface{}) {
	fmt.Println(Cyan("---------------"))
	ps := Sprintf("%s", s)
	fmt.Println(Cyan(ps))
}
