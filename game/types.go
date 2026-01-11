package game

type display struct {
	Buffer []uint32
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
