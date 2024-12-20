package mouse

import (
	"time"
)

type mouse struct {
	x                   float64
	y                   float64
	options             *mouseOptions
	clickHandlers       []func(x, y int)
	holdHandlers        []func(x, y int)
	doubleClickHandlers []func(x, y int)
}

type MouseButton int

const (
	LeftButton MouseButton = iota
	MiddleButton
	RightButton
)

type mouseButton struct {
	button         MouseButton
	mousePressed   bool
	lastClickTime  time.Time
	clickCount     int
	pressStartTime time.Time
	isHolding      bool
}

func NewMouse(options ...MouseOption) *mouse {
	mouseOptions := mouseOptions{}
	for _, option := range options {
		option(&mouseOptions)
	}

	return &mouse{
		options: &mouseOptions,
	}
}

func (m *mouse) UpdateEbiten() {

}
