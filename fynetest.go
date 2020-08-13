package main

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

func main() {
	app := app.New()

	w := app.NewWindow("こんにちは")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("こんにちは Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}