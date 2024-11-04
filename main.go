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
	// Crear la aplicación Fyne
	myApp := app.New()
	stage := myApp.NewWindow("App - Parking Simulation")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(815, 515))
	stage.SetFixedSize(true)

	// Inicializar la escena
	scene := scenes.NewScene(stage)
	scene.Init()

	// Lista de imágenes de coches
	carImages := []string{
		"./assets/car.png",
		"./assets/car2.png",
		"./assets/car3.png",
		"./assets/car4.png",
		"./assets/car5.png",
		"./assets/car6.png",
		"./assets/car7.png",
	}

	// Crear el botón para agregar un coche
	button := widget.NewButton("Agregar Coche", func() {
		b := models.NewCar() // Crear una nueva instancia de coche

		carView := views.NewCar()                       // Crear una nueva vista para el coche
		carView.AddCar(*scene, carImages[imageCounter]) // Agregar el coche a la escena

		imageCounter = (imageCounter + 1) % len(carImages) // Rotar la imagen del coche

		b.Register(carView) // Registrar la vista como observador
		go b.Run()          // Ejecutar la lógica del coche en una goroutine
	})

	// Ajustar la posición y tamaño del botón
	button.Move(fyne.NewPos(90, 100))
	button.Resize(fyne.NewSize(150, 50))
	scene.AddWidget(button) // Agregar el botón a la escena

	// Mostrar la ventana y ejecutar la aplicación
	stage.ShowAndRun()
}
