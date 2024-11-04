package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type Scene struct {
	scene     fyne.Window
	container *fyne.Container
}

func NewScene(scene fyne.Window) *Scene {
	return &Scene{scene: scene, container: nil}
}

func (s *Scene) Init() {
	rect := canvas.NewImageFromURI(storage.NewFileURI("./assets/goFondo.png"))
	rect.Resize(fyne.NewSize(815, 515))
	rect.Move(fyne.NewPos(0, 0))

	// Colocar el rectángulo de fondo y los widgets dentro de un contenedor
	s.container = container.NewWithoutLayout(rect)
	s.scene.SetContent(s.container)
}

func (s *Scene) AddWidget(widget fyne.Widget) {
	s.container.Add(widget)
	s.container.Refresh()
}

func (s *Scene) AddImage(image *canvas.Image) {
	s.container.Add(image)
	s.container.Refresh()
}
