package main

import (
	"golang-games/PuzzleBlock/gameboard"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/guicontrols"
	"golang-games/PuzzleBlock/titlescreen"
	"math/rand"
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
	window, err := sdl.CreateWindow("Loading...", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(WinWidth), int32(WinHeight), sdl.WINDOW_SHOWN)
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

	// Set Random Seed
	rand.Seed(time.Now().UTC().UnixNano())

	// Initialize GameState
	gameStateTransition := gamestatetransition.NewGameStateTransition(WinWidth, WinHeight, gamestate.StartUp, gamestate.TitleScreen, gamestate.StartUp, 500, renderer)

	// TitleScreen variable
	var t *titlescreen.TitleScreen

	// Initialize gameboard
	g := gameboard.NewGameBoard(WinWidth, WinHeight, WinDepth, gameStateTransition, 19, 10, 7, 12, renderer)

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

		switch gameStateTransition.CurrentGameState {
		case gamestate.StartUp:
			// Initialize titlescreen
			t = titlescreen.NewTitleScreen(WinWidth, WinHeight, WinDepth, gameStateTransition, mouseState, 10, renderer)
			gameStateTransition.TransitioningDown = true
			gameStateTransition.CurrentGameState = gamestate.TitleScreen
			window.SetTitle("PuzzleBlock")
		case gamestate.TitleScreen:
			// Get Mouse Input
			if gameStateTransition.TransitioningDown == false {
				mouseState.Update()
			}

			// Draw titlescreen
			t.Update(elapsedTime)
			t.Draw(renderer)

			// Draw transition
			if gameStateTransition.TransitioningDown == true || gameStateTransition.TransitioningUp == true {
				gameStateTransition.Update(elapsedTime)
				gameStateTransition.Draw(renderer)
			}
		case gamestate.OptionsScreen:
			// Get Mouse Input
			if gameStateTransition.Transitioning == false {
				mouseState.Update()
			}

			// Draw transition
			if gameStateTransition.TransitioningDown == true || gameStateTransition.TransitioningUp == true {
				gameStateTransition.Update(elapsedTime)
				gameStateTransition.Draw(renderer)
			}
		case gamestate.MainGame:
			// Get Keyboard Input
			if gameStateTransition.Transitioning == false {
				getKeyboardState(g)
			}

			// Draw gameboard
			g.Update(elapsedTime)
			g.Draw(renderer)

			// Draw transition
			if gameStateTransition.TransitioningDown == true || gameStateTransition.TransitioningUp == true {
				gameStateTransition.Update(elapsedTime)
				gameStateTransition.Draw(renderer)
			}
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
