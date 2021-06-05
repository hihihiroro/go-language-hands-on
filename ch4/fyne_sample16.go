package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")

	mm := fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				fmt.Println("test")
			}),
			fyne.NewMenuItem("Quit", func() {
				a.Quit()
			}),
		),
	)
	w.SetMainMenu(mm)
	w.SetContent(
		container.NewVBox(
			l,
			widget.NewButton("ok", nil),
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
