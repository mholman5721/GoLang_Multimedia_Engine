package gamestatetransition

import (
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/mathhelper"
	"golang-games/PuzzleBlock/texturedrawing"

	"github.com/veandco/go-sdl2/sdl"
)

// GameStateTransition contains all the data needed to transition from one state to another
type GameStateTransition struct {
	WinWidth          int
	WinHeight         int
	FromState         gamestate.GameState
	ToState           gamestate.GameState
	CurrentGameState  gamestate.GameState
	WipeTex           *texturedrawing.SinglePixelTexture
	TransitioningUp   bool
	TransitioningDown bool
	Transitioning     bool
	TransitionTime    float64
	TransitionTimer   float64
}

// NewGameStateTransition creates a new GameStateTransition struct
func NewGameStateTransition(winWidth, winHeight int, fromstate gamestate.GameState, tostate gamestate.GameState, currentstate gamestate.GameState, transitiontime float64, renderer *sdl.Renderer) *GameStateTransition {

	t := texturedrawing.NewSinglePixelTexture(sdl.Color{R: 0, G: 0, B: 0, A: 255}, sdl.Rect{X: 0, Y: 0, W: int32(winWidth), H: int32(winHeight)}, renderer)

	return &GameStateTransition{winWidth, winHeight, fromstate, tostate, currentstate, t, false, false, false, transitiontime, 0}
}

// Update updates the state transition
func (g *GameStateTransition) Update(time float64) {
	if g.TransitionTimer >= g.TransitionTime && g.TransitioningUp == true {
		g.TransitioningUp = false
		g.TransitioningDown = true
		g.FromState = g.CurrentGameState
		g.CurrentGameState = g.ToState

		g.TransitionTimer = 0
	} else if g.TransitionTimer >= g.TransitionTime && g.TransitioningDown == true {
		g.TransitioningDown = false

		g.TransitionTimer = 0
	} else if g.TransitionTimer < g.TransitionTime && g.TransitioningUp == true {
		scale := mathhelper.ScaleBetween(g.TransitionTimer, 0, 1, 0, g.TransitionTime)

		g.WipeTex.Rect.X = int32(g.WinWidth)/2 - int32(float64(g.WinWidth/2)*scale)
		g.WipeTex.Rect.Y = int32(g.WinHeight)/2 - int32(float64(g.WinHeight/2)*scale)
		g.WipeTex.Rect.W = int32(float64(g.WinWidth) * scale)
		g.WipeTex.Rect.H = int32(float64(g.WinHeight) * scale)

		g.WipeTex.Texture.SetAlphaMod(uint8(255 * scale))

		g.TransitionTimer += time
	} else if g.TransitionTimer < g.TransitionTime && g.TransitioningDown == true {
		scale := mathhelper.ScaleBetween(g.TransitionTimer, 0, 1, 0, g.TransitionTime)

		g.WipeTex.Rect.X = int32(float64(g.WinWidth/2) * scale)
		g.WipeTex.Rect.Y = int32(float64(g.WinHeight/2) * scale)
		g.WipeTex.Rect.W = int32(g.WinWidth) - int32(float64(g.WinWidth)*scale)
		g.WipeTex.Rect.H = int32(g.WinHeight) - int32(float64(g.WinHeight)*scale)

		g.WipeTex.Texture.SetAlphaMod(uint8(255 * scale))

		g.TransitionTimer += time
	}
}

// Draw draws the WipeTex for the transition
func (g *GameStateTransition) Draw(renderer *sdl.Renderer) {
	g.WipeTex.Draw(renderer)
}
