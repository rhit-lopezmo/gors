package game

import "image/color"

const (
	FPS = 60

	DISPLAY_WIDTH  = 800
	DISPLAY_HEIGHT = 600

	AUDIO_SAMPLE_RATE     = 44100
	AUDIO_CHANNELS        = 2
	AUDIO_BUFFER_CAPACITY = (AUDIO_SAMPLE_RATE * AUDIO_CHANNELS) / FPS
)

var displayBuffer = make([]color.RGBA, DISPLAY_WIDTH*DISPLAY_HEIGHT)
var audioBuffer = make([]int16, AUDIO_BUFFER_CAPACITY)

func NewGame() *Game {
	for i := range len(displayBuffer) {
		// fill with red
		setPixelColorFromHex(&displayBuffer[i], 0x000000)
	}

	return &Game{
		Fps: FPS,
		Display: display{
			Buffer: displayBuffer,
			Width:  DISPLAY_WIDTH,
			Height: DISPLAY_HEIGHT,
		},
		Audio: audio{
			Buffer:     audioBuffer,
			SampleRate: AUDIO_SAMPLE_RATE,
			Channels:   AUDIO_CHANNELS,
		},
	}
}

func (game *Game) Update() {
	// update pixels

	// update audio
}

func (game *Game) KeyUp(key int) {
	// TODO: implement
}

func (game *Game) KeyDown(key int) {
	// TODO: implement
}

// setPixelColorFromHex sets the pixel color from a 32-bit hex value
//
// hex must either be a 6-digit hex color (0xRRGGBB) or 8-digit hex color (0xRRGGBBAA)
func setPixelColorFromHex(pixel *color.RGBA, hex uint32) {
	// if no alpha present, default to 100%
	if hex <= 0xFFFFFF {
		pixel.R = uint8(hex >> 16)
		pixel.G = uint8(hex >> 8)
		pixel.B = uint8(hex)
		pixel.A = 0xFF
		return
	}

	// with alpha
	pixel.R = uint8(hex >> 24)
	pixel.G = uint8(hex >> 16)
	pixel.B = uint8(hex >> 8)
	pixel.A = uint8(hex)
}
