package game

const (
	FPS = 60

	DISPLAY_WIDTH  = 800
	DISPLAY_HEIGHT = 600

	AUDIO_SAMPLE_RATE     = 44100
	AUDIO_CHANNELS        = 2
	AUDIO_BUFFER_CAPACITY = (AUDIO_SAMPLE_RATE * AUDIO_CHANNELS) / FPS
)

var displayBuffer = make([]uint32, DISPLAY_WIDTH*DISPLAY_HEIGHT)
var audioBuffer = make([]int16, AUDIO_BUFFER_CAPACITY)

func NewGame() *Game {
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
	// TODO: implement

	// update pixels

	// update audio
}

func (game *Game) KeyUp(key int) {
	// TODO: implement
}

func (game *Game) KeyDown(key int) {
	// TODO: implement
}
