package optionsscreen

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/guicontrols"
	"golang-games/PuzzleBlock/musicplayer"
	"golang-games/PuzzleBlock/soundplayer"
	"golang-games/PuzzleBlock/sprite"
	"math/rand"
	"strconv"
	"vec3"

	"github.com/veandco/go-sdl2/sdl"
)

// OptionsScreen is a struct that contains all the sprite information for the options screen
type OptionsScreen struct {
	CurrentGameState      *gamestatetransition.GameStateTransition
	MouseState            *guicontrols.MouseState
	MusicPlayer           *musicplayer.MusicPlayer
	SoundPlayer           *soundplayer.SoundPlayer
	WinWidth              int
	WinHeight             int
	Background            *sprite.Sprite
	TextFont              *font.TTFFont
	TitleText             *font.TTFString
	PreviousCurrentTune   int
	InGameTuneText        *font.TTFString
	InGameTuneValueText   *font.TTFString
	TuneUpButton          *guicontrols.SpriteButton
	TuneDownButton        *guicontrols.SpriteButton
	SoundVolume           int
	PreviousSoundVolume   int
	SoundVolumeText       *font.TTFString
	SoundVolumeValueText  *font.TTFString
	SoundVolumeUpButton   *guicontrols.SpriteButton
	SoundVolumeDownButton *guicontrols.SpriteButton
	MusicVolume           int
	PreviousMusicVolume   int
	MusicVolumeText       *font.TTFString
	MusicVolumeValueText  *font.TTFString
	MusicVolumeUpButton   *guicontrols.SpriteButton
	MusicVolumeDownButton *guicontrols.SpriteButton
	BackButton            *guicontrols.TextButton
}

// NewOptionsScreen is an options screen constructor
func NewOptionsScreen(winWidth, winHeight, winDepth int, gamestate *gamestatetransition.GameStateTransition, mousestate *guicontrols.MouseState, musicplayer *musicplayer.MusicPlayer, soundplayer *soundplayer.SoundPlayer, renderer *sdl.Renderer) *OptionsScreen {

	o := &OptionsScreen{}

	o.CurrentGameState = gamestate

	o.MouseState = mousestate

	o.MusicPlayer = musicplayer

	o.SoundPlayer = soundplayer

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

	o.PreviousCurrentTune = o.MusicPlayer.CurrentTune

	// Set the tune text
	o.InGameTuneText = font.NewTTFString("In-Game Music",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.35, Y: float32(o.WinHeight) * 0.31, Z: 0},
		o.TextFont,
		renderer)

	// Set the tune value text
	o.InGameTuneValueText = font.NewTTFString("Music "+strconv.Itoa(o.MusicPlayer.CurrentTune),
		font.FontMedium,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.115, Y: float32(o.WinHeight) * 0.33, Z: 0},
		o.TextFont,
		renderer)

	o.TuneUpButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowRight.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.25, Y: float32(o.WinHeight) * 0.31, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

	o.TuneDownButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowLeft.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.05, Y: float32(o.WinHeight) * 0.31, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

	o.SoundVolume = 50
	o.PreviousSoundVolume = o.SoundVolume

	// Set the sound volume text
	o.SoundVolumeText = font.NewTTFString("Sound Volume",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.35, Y: float32(o.WinHeight) * 0.48, Z: 0},
		o.TextFont,
		renderer)

	// Set the sound volume value text
	o.SoundVolumeValueText = font.NewTTFString(strconv.Itoa(o.SoundVolume)+" %",
		font.FontMedium,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.14, Y: float32(o.WinHeight) * 0.50, Z: 0},
		o.TextFont,
		renderer)

	o.SoundVolumeUpButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowRight.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.25, Y: float32(o.WinHeight) * 0.48, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

	o.SoundVolumeDownButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowLeft.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.05, Y: float32(o.WinHeight) * 0.48, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

	o.MusicVolume = 50
	o.PreviousMusicVolume = o.MusicVolume

	// Set the music volume text
	o.MusicVolumeText = font.NewTTFString("Music Volume",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.35, Y: float32(o.WinHeight) * 0.65, Z: 0},
		o.TextFont,
		renderer)

	// Set the music volume value text
	o.MusicVolumeValueText = font.NewTTFString(strconv.Itoa(o.MusicVolume)+" %",
		font.FontMedium,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: float32(o.WinWidth) * 0.14, Y: float32(o.WinHeight) * 0.67, Z: 0},
		o.TextFont,
		renderer)

	o.MusicVolumeUpButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowRight.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.25, Y: float32(o.WinHeight) * 0.65, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

	o.MusicVolumeDownButton = guicontrols.NewSpriteButton(o.WinWidth,
		o.WinHeight,
		"assets/arrowLeft.png",
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: float32(o.WinWidth) * 0.05, Y: float32(o.WinHeight) * 0.65, Z: 0},
		0.1,
		100,
		64,
		64,
		1,
		1,
		renderer)

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
		o.MusicPlayer.FutureTune = 0
		o.MusicPlayer.PastTune = o.MusicPlayer.CurrentTune
		o.CurrentGameState.TransitioningUp = true
		o.CurrentGameState.ToState = gamestate.TitleScreen
	}

	// Set the appropriate tune
	if o.TuneUpButton.WasLeftClicked == true {
		o.MusicPlayer.CurrentTune++
		if o.MusicPlayer.CurrentTune > o.MusicPlayer.NumTunes-1 {
			o.MusicPlayer.CurrentTune = 0
		}
		o.MusicPlayer.PlayTune(o.MusicPlayer.CurrentTune)
	}

	if o.TuneDownButton.WasLeftClicked == true {
		o.MusicPlayer.CurrentTune--
		if o.MusicPlayer.CurrentTune < 0 {
			o.MusicPlayer.CurrentTune = o.MusicPlayer.NumTunes - 1
		}
		o.MusicPlayer.PlayTune(o.MusicPlayer.CurrentTune)
	}

	// Set sound volume when appropriate button is clicked
	if o.SoundVolumeUpButton.WasLeftClicked == true {
		o.SoundVolume += 10
		if o.SoundVolume > 100 {
			o.SoundVolume = 100
		}
		o.SoundPlayer.SetVolume(o.SoundVolume)
		o.SoundPlayer.PlaySound("break" + strconv.Itoa(1+rand.Intn(5)))
	}

	if o.SoundVolumeDownButton.WasLeftClicked == true {
		o.SoundVolume -= 10
		if o.SoundVolume < 0 {
			o.SoundVolume = 0
		}
		o.SoundPlayer.SetVolume(o.SoundVolume)
		o.SoundPlayer.PlaySound("break" + strconv.Itoa(1+rand.Intn(5)))
	}

	// Set music volume when appropriate button is clicked
	if o.MusicVolumeUpButton.WasLeftClicked == true {
		o.MusicVolume += 10
		if o.MusicVolume > 100 {
			o.MusicVolume = 100
		}
		o.MusicPlayer.SetVolume(o.MusicVolume)
	}

	if o.MusicVolumeDownButton.WasLeftClicked == true {
		o.MusicVolume -= 10
		if o.MusicVolume < 0 {
			o.MusicVolume = 0
		}
		o.MusicPlayer.SetVolume(o.MusicVolume)
	}

	// Update the buttons
	o.BackButton.Update(o.MouseState, time)
	o.TuneUpButton.Update(o.MouseState, time)
	o.TuneDownButton.Update(o.MouseState, time)
	o.SoundVolumeUpButton.Update(o.MouseState, time)
	o.SoundVolumeDownButton.Update(o.MouseState, time)
	o.MusicVolumeUpButton.Update(o.MouseState, time)
	o.MusicVolumeDownButton.Update(o.MouseState, time)
}

// Draw draws all the objects on the title screen
func (o *OptionsScreen) Draw(renderer *sdl.Renderer) {

	// Draw the background
	o.Background.Draw(renderer)

	// Change the display text depending on whether the underlying value has changed
	if o.MusicPlayer.CurrentTune != o.PreviousCurrentTune {
		o.InGameTuneValueText.ChangeStringTexture("Music "+strconv.Itoa(o.MusicPlayer.CurrentTune), font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		o.PreviousCurrentTune = o.MusicPlayer.CurrentTune
	}

	if o.SoundVolume != o.PreviousSoundVolume {
		if o.SoundVolume == 100 {
			o.SoundVolumeValueText.ChangeStringTexture(strconv.Itoa(o.SoundVolume)+"%", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		} else if o.SoundVolume >= 10 && o.SoundVolume < 100 {
			o.SoundVolumeValueText.ChangeStringTexture(strconv.Itoa(o.SoundVolume)+" %", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		} else if o.SoundVolume < 10 {
			o.SoundVolumeValueText.ChangeStringTexture(strconv.Itoa(o.SoundVolume)+"  %", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		}
		o.PreviousSoundVolume = o.SoundVolume
	}

	if o.MusicVolume != o.PreviousMusicVolume {
		if o.MusicVolume == 100 {
			o.MusicVolumeValueText.ChangeStringTexture(strconv.Itoa(o.MusicVolume)+"%", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		} else if o.MusicVolume >= 10 && o.MusicVolume < 100 {
			o.MusicVolumeValueText.ChangeStringTexture(strconv.Itoa(o.MusicVolume)+" %", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		} else if o.MusicVolume < 10 {
			o.MusicVolumeValueText.ChangeStringTexture(strconv.Itoa(o.MusicVolume)+"  %", font.FontMedium, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		}
		o.PreviousMusicVolume = o.MusicVolume
	}

	// Draw the text
	o.TitleText.Draw(renderer)
	o.InGameTuneText.Draw(renderer)
	o.InGameTuneValueText.Draw(renderer)
	o.SoundVolumeText.Draw(renderer)
	o.SoundVolumeValueText.Draw(renderer)
	o.MusicVolumeText.Draw(renderer)
	o.MusicVolumeValueText.Draw(renderer)

	// Draw the buttons
	o.BackButton.Draw(renderer)
	o.TuneUpButton.Draw(renderer)
	o.TuneDownButton.Draw(renderer)
	o.SoundVolumeUpButton.Draw(renderer)
	o.SoundVolumeDownButton.Draw(renderer)
	o.MusicVolumeUpButton.Draw(renderer)
	o.MusicVolumeDownButton.Draw(renderer)
}
