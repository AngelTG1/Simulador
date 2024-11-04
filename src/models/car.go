package models

import (
	"sync"
	"time"
)

type CarState string

const (
	StateCreated    CarState = "creado"
	StateGate       CarState = "puerta"
	StateParkingBox CarState = "caja"
	StateParked     CarState = "estacionado"
	StateExit       CarState = "salida"
	StateFinished   CarState = "terminado"
)

type Car struct {
	posX, posY int32
	status     bool
	state      CarState
	observers  []Observer
	parkingIdx int
	rotation   float32
}

type Rectangle struct {
	x, y, width, height int32
	occupied            bool
	parkDuration        time.Duration
}

var (
	parkingSpaces = []Rectangle{
		{715, 15, 70, 90, false, 12 * time.Second},
		{715, 75, 70, 90, false, 5 * time.Second},
		{715, 135, 70, 90, false, 7 * time.Second},
		{715, 195, 70, 90, false, 3 * time.Second},
		{715, 255, 70, 90, false, 3 * time.Second},
		{715, 325, 70, 90, false, 4 * time.Second},
		{715, 388, 70, 90, false, 5 * time.Second},
		{715, 450, 70, 90, false, 3 * time.Second},
	}
	parkingMutex sync.Mutex // Mutex para proteger acceso a parkingSpaces
)

// Crea una nueva instancia de Car
func NewCar() *Car {
	return &Car{posX: 6, posY: 255, status: true, state: StateCreated, parkingIdx: -1}
}

// Verifica y reserva un espacio de estacionamiento disponible
func (c *Car) checkParkingSpace() (bool, Rectangle) {
	parkingMutex.Lock()
	defer parkingMutex.Unlock()

	for i, space := range parkingSpaces {
		if !space.occupied {
			parkingSpaces[i].occupied = true
			c.parkingIdx = i
			return true, space
		}
	}
	return false, Rectangle{}
}

// Mueve el auto hacia una posición objetivo de manera más fluida
func (c *Car) moveTo(targetX, targetY int32, step int32, delay time.Duration) {
	// Movimiento en X
	for c.posX != targetX {
		if c.posX < targetX {
			c.posX += step
			if c.posX > targetX {
				c.posX = targetX
			}
		} else {
			c.posX -= step
			if c.posX < targetX {
				c.posX = targetX
			}
		}
		c.NotifyAll()
		time.Sleep(delay)
	}
	// Movimiento en Y
	for c.posY != targetY {
		if c.posY < targetY {
			c.posY += step
			if c.posY > targetY {
				c.posY = targetY
			}
		} else {
			c.posY -= step
			if c.posY < targetY {
				c.posY = targetY
			}
		}
		c.NotifyAll()
		time.Sleep(delay)
	}
}

func (c *Car) Run() {
	var incX int32 = 5
	c.status = true

	for c.status {
		switch c.state {
		case StateCreated:
			c.state = StateGate

		case StateGate:
			c.state = StateParkingBox

		case StateParkingBox:
			if c.posX >= 512 {
				if available, space := c.checkParkingSpace(); available {

					c.moveTo(space.x, space.y, 2, 16*time.Millisecond)
					c.state = StateParked
					c.NotifyAll()
				}
			} else {

				c.posX += incX
				c.NotifyAll()
				time.Sleep(16 * time.Millisecond)
			}

		case StateParked:

			time.Sleep(parkingSpaces[c.parkingIdx].parkDuration)

			// Libera el espacio de estacionamiento
			parkingMutex.Lock()
			parkingSpaces[c.parkingIdx].occupied = false
			parkingMutex.Unlock()

			// Prepara para salir
			c.parkingIdx = -1
			c.state = StateExit
			c.NotifyAll()

		case StateExit:
			// Movimiento suave hacia la salida
			c.moveTo(60, c.posY, 2, 16*time.Millisecond)
			time.Sleep(500 * time.Millisecond)

			c.rotation += 90
			c.NotifyAll()

			c.moveTo(6, 175, 2, 16*time.Millisecond)
			c.state = StateFinished

		case StateFinished:
			c.status = false
		}

		time.Sleep(16 * time.Millisecond)
	}
}

func (c *Car) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *Car) NotifyAll() {
	for _, observer := range c.observers {
		observer.Update(Pos{X: c.posX, Y: c.posY, Rotation: c.rotation})
	}
}
