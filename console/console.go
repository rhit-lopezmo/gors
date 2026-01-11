package console

import (
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
	gorsGame "github.com/rhit-lopezmo/gors/game"
)

var tex *rl.Texture2D
var img *rl.Image

func Configure(game *gorsGame.Game) {
	rl.SetTraceLogLevel(rl.LogWarning)
	rl.SetTargetFPS(int32(game.Fps))

	rl.InitWindow(int32(game.Display.Width), int32(game.Display.Height), "Game")
	rl.SetWindowSize(game.Display.Width, game.Display.Height)
	rl.SetWindowMinSize(game.Display.Width, game.Display.Height)
}

func InitMainTexture(game *gorsGame.Game) *rl.Texture2D {
	displayBufferPtr := unsafe.Pointer(&game.Display.Buffer[0])
	displayBufferAsBytes := unsafe.Slice(
		(*byte)(displayBufferPtr),
		uintptr(len(game.Display.Buffer))*unsafe.Sizeof(game.Display.Buffer[0]),
	)

	img = rl.NewImage(
		displayBufferAsBytes,
		int32(game.Display.Width),
		int32(game.Display.Height),
		1,
		rl.UncompressedR8g8b8a8,
	)

	loadedTex := rl.LoadTextureFromImage(img)
	tex = &loadedTex

	return tex
}

func StartGame(game *gorsGame.Game, tex *rl.Texture2D) {
	for !rl.WindowShouldClose() {
		game.Update()
		rl.UpdateTexture(*tex, game.Display.Buffer)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTexturePro(
			*tex,
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

func CleanUp() {
	rl.UnloadImage(img)
	rl.UnloadTexture(*tex)
	rl.CloseWindow()
}
