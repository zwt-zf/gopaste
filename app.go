package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/zwt-zf/gopaste/lib"
)

func main() {

	// addClipboardFormatListener()

	a := app.New()
	w := a.NewWindow("Hello")

	w.Resize(fyne.NewSize(500, 500))

	lib.RegisterListener(2)

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
