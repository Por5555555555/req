package colortext

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func IsOk[T any](Gstring T) {
	textColor := color.New(color.FgGreen)
	textOut := color.New(color.FgHiBlue).Add(color.Underline)

	textColor.Print("Ok")
	fmt.Print(" : ")
	textOut.Println(Gstring)
}

func IsErr[T any](Gstring T) {
	textColor := color.New(color.FgRed)
	textOut := color.New(color.FgHiYellow).Add(color.Underline).SprintFunc()

	textColor.Print("Err")
	fmt.Print(" : ")
	log.Fatalln(textOut(Gstring))
}

func IsTime[T any](Gstring T) {
	textColor := color.New(color.FgYellow)
	textOut := color.New(color.FgHiBlack).Add(color.Underline).SprintFunc()

	textColor.Print("Time")
	fmt.Print(" : ")
	fmt.Println(textOut(Gstring))
}

func IsWanging[T any](Gstring T) {
	Wanging := color.New(color.FgYellow).Add(color.Underline).SprintFunc()
	fmt.Println("Wanging", Wanging(Gstring))
}

func IsLogErr[T any](Gstring T) {
	LogErr := color.New(color.FgCyan).Add(color.Underline).SprintFunc()
	fmt.Println("Log Error", LogErr(Gstring))
}
