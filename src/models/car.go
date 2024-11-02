package models

import (
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
}

type Rectangle struct {
	x, y, width, height int32
	occupied            bool
	parkDuration        time.Duration 
}

var parkingSpaces = []Rectangle{
	{715, 15, 70, 90, false, 4 * time.Second},  
	{715, 75, 70, 90, false, 5 * time.Second},   
	{715, 135, 70, 90, false, 5 * time.Second}, 
	{715, 195, 70, 90, false, 3 * time.Second}, 
	{715, 255, 70, 90, false, 2 * time.Second}, 
	{715, 320, 70, 90, false, 5 * time.Second},
	{715, 380, 70, 90, false, 3 * time.Second}, 
	{715, 440, 70, 90, false, 5 * time.Second},  
}

// Crea una nueva instancia de Car
func NewCar() *Car {
	return &Car{posX: 60, posY: 255, status: true, state: StateCreated, parkingIdx: -1}
}

func (c *Car) checkParkingSpace() bool {
	for {
		for i, space := range parkingSpaces {
			if !space.occupied {

				c.posX = space.x
				c.posY = space.y
				parkingSpaces[i].occupied = true
				c.parkingIdx = i
				c.state = StateParked
				return true
			}
		}
		
		time.Sleep(500 * time.Millisecond)
	}
}

func (c *Car) Run() {
	var incX int32 = 30
	c.status = true

	for c.status {
		switch c.state {
		case StateCreated:
			c.state = StateGate

		case StateGate:
			c.state = StateParkingBox

		case StateParkingBox:
			
			if c.checkParkingSpace() {
				c.NotifyAll() 
			} else {
				
				c.posX += incX
				c.NotifyAll()
				time.Sleep(500 * time.Millisecond)
			}

		case StateParked:
			
			time.Sleep(parkingSpaces[c.parkingIdx].parkDuration)
			
			parkingSpaces[c.parkingIdx].occupied = false
			c.parkingIdx = -1
			c.state = StateExit
			c.NotifyAll() 

		case StateExit:
			c.posX -= incX
			c.NotifyAll()
			time.Sleep(500 * time.Millisecond)

			if c.posX < 60 {
				c.state = StateFinished
			}

		case StateFinished:
			c.status = false 
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (c *Car) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *Car) NotifyAll() {
	for _, observer := range c.observers {
		observer.Update(Pos{X: c.posX, Y: c.posY})
	}
}
