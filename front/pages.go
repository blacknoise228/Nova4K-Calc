package front

import (
	"ResolutionCalc/maths"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (app *FyneApp) Mctrl4kPage() {
	w := app.Window
	resolutionForm := widget.NewForm(
		widget.NewFormItem("Width", widget.NewEntry()),
		widget.NewFormItem("Height", widget.NewEntry()),
		widget.NewFormItem("Refresh Rate", widget.NewSelect([]string{"60Hz", "50Hz"}, nil)),
	)
	resolutionForm.OnSubmit = func() {
		ok, err := maths.MCTRL4kResolution(resolutionForm.Items[0].Widget.(*widget.Entry).Text,
			resolutionForm.Items[1].Widget.(*widget.Entry).Text,
			resolutionForm.Items[2].Widget.(*widget.Select).Selected)
		if err != nil {
			dialog.ShowError(err, w)
		}
		if ok {
			dialog.ShowInformation("OK", "Resolution OK", w)
		}
	}
	back := widget.NewButton("Back", func() {
		app.WelcomePage()
	})
	text := widget.NewLabel("MCTRL4K")
	text.Alignment = fyne.TextAlignCenter
	content := container.NewVBox(text, resolutionForm, back)
	w.SetContent(content)
	w.Show()
}
func (app *FyneApp) Vx1000Page() {
	w := app.Window
	resolutionForm := widget.NewForm(
		widget.NewFormItem("Width", widget.NewEntry()),
		widget.NewFormItem("Height", widget.NewEntry()),
		widget.NewFormItem("Refresh Rate", widget.NewSelect([]string{"60Hz", "50Hz", "25Hz"}, nil)),
	)
	resolutionForm.OnSubmit = func() {
		ok, err := maths.VX1000Resolution(resolutionForm.Items[0].Widget.(*widget.Entry).Text,
			resolutionForm.Items[1].Widget.(*widget.Entry).Text,
			resolutionForm.Items[2].Widget.(*widget.Select).Selected)
		if err != nil {
			dialog.ShowError(err, w)
		}
		if ok {
			dialog.ShowInformation("OK", "Resolution OK", w)
		}
	}
	back := widget.NewButton("Back", func() {
		app.WelcomePage()
	})
	text := widget.NewLabel("VX1000")
	text.Alignment = fyne.TextAlignCenter
	content := container.NewVBox(text, resolutionForm, back)
	w.SetContent(content)
	w.Show()
}
func (app *FyneApp) Mctrl660proPage() {
	w := app.Window
	resolutionForm := widget.NewForm(
		widget.NewFormItem("Width", widget.NewEntry()),
		widget.NewFormItem("Height", widget.NewEntry()),
		widget.NewFormItem("Refresh Rate", widget.NewSelect([]string{"60Hz", "50Hz"}, nil)),
	)
	resolutionForm.OnSubmit = func() {
		ok, err := maths.MCTRL660ProResolution(resolutionForm.Items[0].Widget.(*widget.Entry).Text,
			resolutionForm.Items[1].Widget.(*widget.Entry).Text,
			resolutionForm.Items[2].Widget.(*widget.Select).Selected)
		if err != nil {
			dialog.ShowError(err, w)
		}
		if ok {
			dialog.ShowInformation("OK", "Resolution OK", w)
		}
	}
	back := widget.NewButton("Back", func() {
		app.WelcomePage()
	})
	text := widget.NewLabel("MCTRL660Pro")
	text.Alignment = fyne.TextAlignCenter
	content := container.NewVBox(text, resolutionForm, back)
	w.SetContent(content)
	w.Show()
}
