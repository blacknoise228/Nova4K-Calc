package front

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type FyneApp struct {
	App    fyne.App
	Window fyne.Window
}

func NewFyneApp(app fyne.App) *FyneApp {
	w := app.NewWindow("Resolution Calculator")
	return &FyneApp{App: app, Window: w}
}
func (app *FyneApp) StartApp() {
	app.WelcomePage()
	app.Window.ShowAndRun()

}
func (app *FyneApp) WelcomePage() {
	w := app.Window

	hello := widget.NewLabel("Welcome to the Resolution Calculator")
	hello.Alignment = fyne.TextAlignCenter
	hello2 := widget.NewLabel("Select your sending card")
	hello2.Alignment = fyne.TextAlignCenter

	mctrl4k := widget.NewButton("MCTRL4K", func() {
		app.Mctrl4kPage()
	})
	vx1000 := widget.NewButton("VX1000", func() {
		app.Vx1000Page()
	})
	mctrl660pro := widget.NewButton("MCTRL660Pro", func() {
		app.Mctrl660proPage()
	})
	exit := widget.NewButton("Exit", func() {
		app.App.Quit()
	})
	space := widget.NewLabel(" ")
	content := container.NewVBox(hello, hello2, mctrl4k, vx1000, mctrl660pro, space, exit)
	w.SetContent(content)
	w.SetFullScreen(true)
	w.Show()
}
