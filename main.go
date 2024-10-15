package main

import (
	"ResolutionCalc/front"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	app := front.NewFyneApp(a)
	app.StartApp()
}
