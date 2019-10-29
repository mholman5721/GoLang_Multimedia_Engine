package titlescreen

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/guicontrols"
	"golang-games/PuzzleBlock/musicplayer"
	"golang-games/PuzzleBlock/soundplayer"
	"golang-games/PuzzleBlock/sprite"
	"math/rand"
	"vec3"

	"github.com/veandco/go-sdl2/sdl"
)

// TitleScreen is a struct that contains all the sprite information for the title screen
type TitleScreen struct {
	CurrentGameState *gamestatetransition.GameStateTransition
	MouseState       *guicontrols.MouseState
	MusicPlayer      *musicplayer.MusicPlayer
	SoundPlayer      *soundplayer.SoundPlayer
	WinWidth         int
	WinHeight        int
	Blocks           []*sprite.Sprite
	Background       *sprite.Sprite
	TextFont         *font.TTFFont
	TitleText        *font.TTFString
	StartButton      *guicontrols.TextButton
	OptionsButton    *guicontrols.TextButton
	QuitButton       *guicontrols.TextButton
}

// NewTitleScreen is a title screen constructor
func NewTitleScreen(winWidth, winHeight, winDepth, numBlocks int, gamestate *gamestatetransition.GameStateTransition, mousestate *guicontrols.MouseState, musicplayer *musicplayer.MusicPlayer, soundplayer *soundplayer.SoundPlayer, renderer *sdl.Renderer) *TitleScreen {

	t := &TitleScreen{}

	t.CurrentGameState = gamestate

	t.MouseState = mousestate

	t.MusicPlayer = musicplayer

	t.SoundPlayer = soundplayer

	t.WinWidth = winWidth
	t.WinHeight = winHeight

	t.Blocks = make([]*sprite.Sprite, numBlocks)

	// Set the floating blocks
	for i := 0; i < numBlocks; i++ {
		blockScale := 3 * rand.Float64()
		scaledBlockPixSize := 64 * blockScale
		t.Blocks[i] = sprite.NewSprite(
			"assets/Gems.png",
			vec3.Vector3{
				X: float32(rand.Intn(winWidth - int(scaledBlockPixSize))),
				Y: float32(rand.Intn(winHeight - int(scaledBlockPixSize))),
				Z: float32(winDepth)},
			vec3.Vector3{
				X: (float32(rand.Intn(10)) - float32(rand.Intn(10))*2) * 0.1,
				Y: (float32(rand.Intn(10)) - float32(rand.Intn(10))*2) * 0.1,
				Z: 0},
			64,
			64,
			blockScale,
			blockScale,
			10,
			7,
			rand.Intn(10),
			rand.Intn(7),
			true,
			100,
			true,
			renderer)
		t.Blocks[i].SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 64})
	}

	// Set the background image
	t.Background = sprite.NewSprite(
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
	t.TextFont = font.NewTTFFont("assets/FifteenTwenty-Bold.otf", winWidth, winHeight)

	// Set the title text
	t.TitleText = font.NewTTFString("PuzzleBlock",
		font.FontTitle,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: 0, Y: float32(winHeight) * 0.05, Z: 0},
		t.TextFont,
		renderer)
	t.TitleText.SetCenterX()

	t.StartButton = guicontrols.NewTextButton(t.WinWidth,
		t.WinHeight,
		"   Start   ",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: 0, Y: float32(t.WinHeight) * 0.48, Z: 0},
		0.1,
		100,
		t.TextFont,
		renderer)
	t.StartButton.SetCenterX()

	t.OptionsButton = guicontrols.NewTextButton(t.WinWidth,
		t.WinHeight,
		"  Options  ",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: 0, Y: float32(t.WinHeight) * 0.65, Z: 0},
		0.1,
		100,
		t.TextFont,
		renderer)
	t.OptionsButton.SetCenterX()

	t.QuitButton = guicontrols.NewTextButton(t.WinWidth,
		t.WinHeight,
		"   Quit!   ",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		sdl.Color{R: 128, G: 128, B: 128, A: 192},
		sdl.Color{R: 128, G: 128, B: 192, A: 192},
		sdl.Color{R: 0, G: 0, B: 255, A: 192},
		vec3.Vector3{X: 0, Y: float32(t.WinHeight) * 0.82, Z: 0},
		0.1,
		100,
		t.TextFont,
		renderer)
	t.QuitButton.SetCenterX()

	return t
}

// Update updates all the objects on the title screen
func (t *TitleScreen) Update(time float64) {

	// Change to MainGame if the start button is clicked
	if t.StartButton.WasLeftClicked == true {
		t.MusicPlayer.FutureTune = t.MusicPlayer.PastTune
		t.CurrentGameState.TransitioningUp = true
		t.CurrentGameState.ToState = gamestate.MainGame
	}

	// Change to Options screen if the start button is clicked
	if t.OptionsButton.WasLeftClicked == true {
		t.MusicPlayer.FutureTune = t.MusicPlayer.PastTune
		t.CurrentGameState.TransitioningUp = true
		t.CurrentGameState.ToState = gamestate.OptionsScreen
	}

	// Quit the game if the quit button is clicked
	if t.QuitButton.WasLeftClicked == true {
		t.CurrentGameState.TransitioningUp = true
		t.CurrentGameState.ToState = gamestate.QuitGame
	}

	// Update the blocks
	for i := range t.Blocks {
		if t.Blocks[i].Pos.X < 0 {
			t.Blocks[i].Pos.X = 0
			t.Blocks[i].Vel.X *= -1
		} else if t.Blocks[i].Pos.X > float32(t.WinWidth-int(float64(t.Blocks[i].W)*t.Blocks[i].ScaleX)) {
			t.Blocks[i].Pos.X = float32(t.WinWidth - int(float64(t.Blocks[i].W)*t.Blocks[i].ScaleX))
			t.Blocks[i].Vel.X *= -1
		}
		if t.Blocks[i].Pos.Y < 0 {
			t.Blocks[i].Pos.Y = 0
			t.Blocks[i].Vel.Y *= -1
		} else if t.Blocks[i].Pos.Y > float32(t.WinHeight-int(float64(t.Blocks[i].H)*t.Blocks[i].ScaleY)) {
			t.Blocks[i].Pos.Y = float32(t.WinHeight - int(float64(t.Blocks[i].H)*t.Blocks[i].ScaleY))
			t.Blocks[i].Vel.Y *= -1
		}
		t.Blocks[i].Update(time)
	}

	// Update the buttons
	t.StartButton.Update(t.MouseState, time)
	t.OptionsButton.Update(t.MouseState, time)
	t.QuitButton.Update(t.MouseState, time)
}

// Draw draws all the objects on the title screen
func (t *TitleScreen) Draw(renderer *sdl.Renderer) {

	// Draw the background
	t.Background.Draw(renderer)

	// Draw the blocks
	for i := range t.Blocks {
		t.Blocks[i].Draw(renderer)
	}

	// Draw the text
	t.TitleText.Draw(renderer)

	// Draw the buttons
	t.StartButton.Draw(renderer)
	t.OptionsButton.Draw(renderer)
	t.QuitButton.Draw(renderer)
}
