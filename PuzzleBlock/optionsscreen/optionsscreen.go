package optionsscreen

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/guicontrols"
	"golang-games/PuzzleBlock/sprite"
	"vec3"

	"github.com/veandco/go-sdl2/sdl"
)

// OptionsScreen is a struct that contains all the sprite information for the options screen
type OptionsScreen struct {
	CurrentGameState *gamestatetransition.GameStateTransition
	MouseState       *guicontrols.MouseState
	WinWidth         int
	WinHeight        int
	Background       *sprite.Sprite
	TextFont         *font.TTFFont
	TitleText        *font.TTFString
	/*SoundVolumeUpButton   *guicontrols.TextButton
	SoundVolumeDownButton *guicontrols.TextButton
	MusicVolumeUpButton   *guicontrols.TextButton
	MusicVolumeDownButton *guicontrols.TextButton*/
	BackButton *guicontrols.TextButton
}

// NewOptionsScreen is an options screen constructor
func NewOptionsScreen(winWidth, winHeight, winDepth int, gamestate *gamestatetransition.GameStateTransition, mousestate *guicontrols.MouseState, renderer *sdl.Renderer) *OptionsScreen {

	o := &OptionsScreen{}

	o.CurrentGameState = gamestate

	o.MouseState = mousestate

	o.WinWidth = winWidth
	o.WinHeight = winHeight

	// Set the background image
	o.Background = sprite.NewSprite(
		"assets/background.png",
		vec3.Vector3{X: 0, Y: 0, Z: 0},
		vec3.Vector3{X: 0, Y: 0, Z: 0},
		1280,
		720,
		float64(winWidth)/1280,
		float64(winHeight)/720,
		1,
		1,
		0,
		0,
		true,
		0,
		false,
		renderer)

	// Set the font for the text
	o.TextFont = font.NewTTFFont("assets/FifteenTwenty-Bold.otf", winWidth, winHeight)

	// Set the title text
	o.TitleText = font.NewTTFString("Options",
		font.FontTitle,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: 0, Y: float32(winHeight) * 0.05, Z: 0},
		o.TextFont,
		renderer)
	o.TitleText.SetCenterX()

	o.BackButton = guicontrols.NewTextButton(o.WinWidth,
		o.WinHeight,
		"   Back   ",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: 0, Y: float32(o.WinHeight) * 0.82, Z: 0},
		0.1,
		100,
		o.TextFont,
		renderer)
	o.BackButton.SetCenterX()

	return o
}

// Update updates all the objects on the title screen
func (o *OptionsScreen) Update(time float64) {

	// Return to the title screen if back button is clicked
	if o.BackButton.WasLeftClicked == true {
		o.CurrentGameState.TransitioningUp = true
		o.CurrentGameState.ToState = gamestate.TitleScreen
	}

	// Update the buttons
	o.BackButton.Update(o.MouseState, time)
}

// Draw draws all the objects on the title screen
func (o *OptionsScreen) Draw(renderer *sdl.Renderer) {

	// Draw the background
	o.Background.Draw(renderer)

	// Draw the text
	o.TitleText.Draw(renderer)

	// Draw the buttons
	o.BackButton.Draw(renderer)
}
