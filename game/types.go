package game

import "image/color"

type display struct {
	Buffer []color.RGBA
	Width  int
	Height int
}

type audio struct {
	Buffer     []int16
	SampleRate int
	Channels   int
}

type Game struct {
	Fps     int
	Display display
	Audio   audio
}
