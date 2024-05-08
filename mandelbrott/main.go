package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	helloLabel := widget.NewLabel("Hello Fyne!")
	helloButton := widget.NewButton("Click me!", func() {
		helloLabel.SetText("Button Clicked!")
	})

	myWindow.SetContent(container.NewVBox(
		helloLabel,
		helloButton,
	))

	myWindow.ShowAndRun()
}
