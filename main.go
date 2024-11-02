package main

import (
	"fynego/src/models"
	"fynego/src/scenes"
	"fynego/src/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// Variable para controlar la rotación de imágenes
var imageCounter int

func main() {
	myApp := app.New()
	stage := myApp.NewWindow("App - Parking Simulation")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(815, 515))
	stage.SetFixedSize(true)

	scene := scenes.NewScene(stage)
	scene.Init()

	carImages := []string{
		"./assets/car.png",
		"./assets/car2.png",
		"./assets/car3.png",
		"./assets/car4.png",
		"./assets/car5.png",
		"./assets/car6.png",
		"./assets/car7.png",
	}

	button := widget.NewButton("Agregar Coche", func() {
		b := models.NewCar()

		carView := views.NewCar()
		carView.AddCar(*scene, carImages[imageCounter])

		imageCounter = (imageCounter + 1) % len(carImages)

		b.Register(carView)
		go b.Run()
	})

	button.Move(fyne.NewPos(90, 100))
	button.Resize(fyne.NewSize(150, 50))
	scene.AddWidget(button)

	stage.ShowAndRun()
}
