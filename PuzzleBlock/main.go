package main

import (
	"golang-games/PuzzleBlock/gameboard"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/guicontrols"
	"golang-games/PuzzleBlock/titlescreen"
	"time"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func init() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	err = mix.Init(mix.INIT_OGG)
	if err != nil {
		panic(err)
	}
}

func initRendererAndWindow() (*sdl.Renderer, *sdl.Window) {
	window, err := sdl.CreateWindow("PuzzleBlock", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(WinWidth), int32(WinHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	return renderer, window
}

func main() {

	// Timing variables
	var frameStart time.Time
	var elapsedTime float64

	// Initialize renderer
	renderer, window = initRendererAndWindow()

	// Initialize input
	initInput()

	mouseState := guicontrols.GetMouseState()

	// Initialize GameState
	currentGameState := gamestate.TitleScreen

	// Initialize titlescreen
	t := titlescreen.NewTitleScreen(&currentGameState, mouseState, WinWidth, WinHeight, WinDepth, 10, renderer)

	// Initialize gameboard
	g := gameboard.NewGameBoard(&currentGameState, WinWidth, WinHeight, WinDepth, 19, 10, 7, 12, renderer)

	// Main game loop
	for {
		frameStart = time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.TouchFingerEvent:
				if e.Type == sdl.FINGERDOWN {
					//touchX := int(e.X * float32(WinWidth))
					//touchY := int(e.Y * float32(WinHeight))
					//currentMouseState.x = touchX
					//currentMouseState.y = touchY
					//currentMouseState.leftButton = true
				}
			}
		}

		// Clear the screen
		renderer.Clear()

		switch currentGameState {
		case gamestate.TitleScreen:
			// Get Mouse Input
			mouseState.Update()

			// Draw titlescreen
			t.Update(elapsedTime)
			t.Draw(renderer)
		case gamestate.OptionsScreen:
			// Get Mouse Input
			mouseState.Update()
		case gamestate.MainGame:
			// Get Keyboard Input
			getKeyboardState(g)

			// Draw gameboard
			g.Update(elapsedTime)
			g.Draw(renderer)
		case gamestate.QuitGame:
			return
		default:
		}

		// Update Window Texture
		renderer.Present()

		elapsedTime = time.Since(frameStart).Seconds() * 1000

		if elapsedTime < 5 {
			sdl.Delay(5 - uint32(elapsedTime))
			elapsedTime = time.Since(frameStart).Seconds() * 1000
			//fps := int(1000 / elapsedTime)
			//fmt.Println("ms per frame:", elapsedTime, "|", "fps:", fps)
		}
	}
}
