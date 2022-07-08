package rdisplay

import "image"

// ScreenGrabberInterface
type ScreenGrabberInterface interface {
	Start()
	Frames() <-chan *image.RGBA
	Stop()
	Fps() int
	Screen() *Screen
}

// Screen
type Screen struct {
	Index  int
	Bounds image.Rectangle
}

// Service
type ServiceInterface interface {
	CreateScreenGrabber(screen Screen, fps int) (ScreenGrabberInterface, error)
	Screens() ([]Screen, error)
}
