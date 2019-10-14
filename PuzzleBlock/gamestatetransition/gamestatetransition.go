package gamestatetransition

import (
	"golang-games/PuzzleBlock/gamestate"
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
func NewGameStateTransition(winWidth, winHeight int, fromstate gamestate.GameState, tostate, currentstate gamestate.GameState, transitiontime float64, renderer *sdl.Renderer) *GameStateTransition {

	t := texturedrawing.NewSinglePixelTexture(sdl.Color{R: 0, G: 0, B: 0, A: 0}, sdl.Rect{X: int32(winWidth / 2), Y: int32(winHeight / 2), W: 1, H: 1}, renderer)

	return &GameStateTransition{winWidth, winHeight, fromstate, tostate, currentstate, t, false, false, false, transitiontime, 0}
}

// Update updates the state transition
func (g *GameStateTransition) Update(time float64) {
	if g.TransitionTimer >= g.TransitionTime {
		g.Transitioning = false
		g.FromState = g.CurrentGameState
		g.CurrentGameState = g.ToState
		/*switch g.CurrentGameState {
		case gamestate.StartUp:
			g.ToState = gamestate.TitleScreen
		case gamestate.TitleScreen:
			g.ToState = gamestate.MainGame
		case gamestate.MainGame:
			g.ToState = gamestate.TitleScreen
		case gamestate.OptionsScreen:
			g.ToState = gamestate.TitleScreen
		case gamestate.QuitGame:
		default:
		}*/
		g.TransitionTimer = 0
	} else {
		g.TransitionTimer += time
	}
}

// Draw draws the WipeTex for the transition
func (g *GameStateTransition) Draw(renderer *sdl.Renderer) {
	g.WipeTex.Draw(renderer)
}
