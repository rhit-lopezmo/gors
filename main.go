package main

import (
	"fmt"

	"github.com/rhit-lopezmo/gors/console"
	gorsGame "github.com/rhit-lopezmo/gors/game"
)

func main() {
	game := gorsGame.NewGame()

	fmt.Println("Game Info:")

	fmt.Println("FPS:", game.Fps)

	fmt.Println("Display Width:", game.Display.Width)
	fmt.Println("Display Height:", game.Display.Height)

	fmt.Println("Audio Sample Rate:", game.Audio.SampleRate)
	fmt.Println("Audio Channels:", game.Audio.Channels)

	console.Configure(game)
	tex := console.InitMainTexture(game)
	defer console.CleanUp()

	console.StartGame(game, tex)
}
