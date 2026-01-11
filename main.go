package main

import (
	"fmt"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(int32(game.Display.Width), int32(game.Display.Height), "Game")
	rl.SetWindowSize(game.Display.Width, game.Display.Height)
	rl.SetWindowMinSize(game.Display.Width, game.Display.Height)
	defer rl.CloseWindow()

	rl.SetTargetFPS(int32(game.Fps))

	displayBufferPtr := unsafe.Pointer(&game.Display.Buffer[0])
	displayBufferAsBytes := unsafe.Slice(
		(*byte)(displayBufferPtr),
		uintptr(len(game.Display.Buffer))*unsafe.Sizeof(game.Display.Buffer[0]),
	)

	img := rl.NewImage(
		displayBufferAsBytes,
		int32(game.Display.Width),
		int32(game.Display.Height),
		1,
		rl.UncompressedR8g8b8a8,
	)
	defer rl.UnloadImage(img)

	tex := rl.LoadTextureFromImage(img)
	defer rl.UnloadTexture(tex)

	for !rl.WindowShouldClose() {
		game.Update()
		rl.UpdateTexture(tex, game.Display.Buffer)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTexturePro(
			tex,
			// SOURCE: entire texture (texture space)
			rl.NewRectangle(
				0,
				0,
				float32(tex.Width),
				float32(tex.Height),
			),
			// DEST: entire window (screen space)
			rl.NewRectangle(
				0,
				0,
				float32(game.Display.Width),
				float32(game.Display.Height),
			),
			// ORIGIN: top-left
			rl.NewVector2(0, 0),
			// ROTATION
			0,
			// TINT
			rl.White,
		)

		rl.EndDrawing()
	}
}
