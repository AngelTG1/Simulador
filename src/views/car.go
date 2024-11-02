package views

import (
	"fmt"
	"fynego/src/models"
	"fynego/src/scenes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Car struct {
	Car *canvas.Image
}

func NewCar() *Car {
	return &Car{Car: nil}
}

func (b *Car) AddCar(c scenes.Scene, imagePath string) {
	carro := canvas.NewImageFromURI(storage.NewFileURI(imagePath))
	carro.Resize(fyne.NewSize(60, 40))
	carro.Move(fyne.NewPos(205, 205))
	b.Car = carro
	c.AddImage(carro)
}

func (b *Car) Update(pos models.Pos) {
	fmt.Printf("Posición del coche: %d : %d\n", pos.X, pos.Y)
	b.Car.Move(fyne.NewPos(float32(pos.X), float32(pos.Y)))
}
