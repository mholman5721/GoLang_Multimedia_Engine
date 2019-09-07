package main

import (
	"golang-games/PuzzleBlock/gameboard"
	"golang-games/PuzzleBlock/guicontrols"
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

	// Initialize GameState
	//var currentGameState GameState
	currentGameState := MainGame

	// Initialize gameboard
	g := gameboard.NewGameBoard(WinWidth, WinHeight, WinDepth, 19, 10, 7, 12, renderer)

	// Test Buttons
	/*
		textButtonFont := font.NewTTFFont("assets/FifteenTwenty-Bold.otf", WinWidth)
		textButton := guicontrols.NewTextButton("  Hello  ",
			font.FontLarge,
			sdl.Color{R: 255, G: 255, B: 255, A: 255},
			sdl.Color{R: 128, G: 128, B: 128, A: 192},
			sdl.Color{R: 128, G: 128, B: 192, A: 192},
			sdl.Color{R: 0, G: 0, B: 255, A: 192},
			vec3.Vector3{X: 200, Y: 500, Z: 0},
			0.3,
			50,
			textButtonFont,
			renderer)

		spriteButton := guicontrols.NewSpriteButton("assets/arrowLeft.png",
			sdl.Color{R: 128, G: 128, B: 128, A: 192},
			sdl.Color{R: 128, G: 128, B: 192, A: 192},
			sdl.Color{R: 0, G: 0, B: 255, A: 192},
			vec3.Vector3{X: 800, Y: 500, Z: 0},
			0.3,
			50,
			64,
			64,
			1,
			1,
			renderer)
	*/

	mouseState := guicontrols.GetMouseState()

	// Main game loop
	for {
		frameStart = time.Now()

		// Input
		/*
			currentMouseState := getMouseState()
			if prevMouseState.leftButton == true {
				fmt.Println("Mouse Down!")
			}
		*/

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.TouchFingerEvent:
				if e.Type == sdl.FINGERDOWN {
					touchX := int(e.X * float32(WinWidth))
					touchY := int(e.Y * float32(WinHeight))
					currentMouseState.x = touchX
					currentMouseState.y = touchY
					currentMouseState.leftButton = true
				}
			}
		}

		// Clear the screen
		renderer.Clear()

		switch currentGameState {
		case TitleScreen:
			// Get Mouse Input
			mouseState.Update()
		case OptionsScreen:
			// Get Mouse Input
			mouseState.Update()
		case MainGame:
			// Get Keyboard Input
			getKeyboardState(g)
			// Draw gameboard
			g.Update(elapsedTime)
			g.Draw(renderer)
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

		prevMouseState = currentMouseState
	}
}
