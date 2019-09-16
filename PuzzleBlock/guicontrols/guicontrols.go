package guicontrols

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/sprite"
	"golang-games/PuzzleBlock/texturedrawing"
	"vec3"

	"github.com/veandco/go-sdl2/sdl"
)

// MouseState structs contain all the information provided by the mouse
type MouseState struct {
	LeftButton      bool
	RightButton     bool
	PrevLeftButton  bool
	PrevRightButton bool
	PrevX, PrevY    int
	X, Y            int
}

// GetMouseState returns a pointer to a MouseState struct with the current mouse information
func GetMouseState() *MouseState {
	mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()

	var result MouseState

	result.X = int(mouseX)
	result.Y = int(mouseY)
	result.LeftButton = !(leftButton == 0)
	result.RightButton = !(rightButton == 0)

	return &result
}

// Update updates the mouse information every 'frame'
func (mouseState *MouseState) Update() {
	mouseState.PrevX = mouseState.X
	mouseState.PrevY = mouseState.Y
	mouseState.PrevLeftButton = mouseState.LeftButton
	mouseState.PrevRightButton = mouseState.RightButton

	X, Y, mouseButtonState := sdl.GetMouseState()

	mouseState.X = int(X)
	mouseState.Y = int(Y)
	mouseState.LeftButton = !((mouseButtonState & sdl.ButtonLMask()) == 0)
	mouseState.RightButton = !((mouseButtonState & sdl.ButtonRMask()) == 0)
}

// SpriteButton structs contain all the information needed for a simple button with an image
type SpriteButton struct {
	MainSprite      *sprite.Sprite
	Rect            sdl.Rect
	Background      *sdl.Texture
	AnimBackground  *sdl.Texture
	SelectedTex     *sdl.Texture
	IsSelected      bool
	WasLeftClicked  bool
	WasRightClicked bool
	AnimSpeed       int
	Animating       bool
	AnimTimer       float64
	BackgroundPos   vec3.Vector3
	SpritePos       vec3.Vector3
	W               int32
	H               int32
}

// NewSpriteButton is a 'constructor' for an SpriteButton struct
func NewSpriteButton(path string, backgroundColor, animBackgroundColor, selectedColor sdl.Color, pos vec3.Vector3, borderPct float32, animSpeedMS, w, h int, scaleX, scaleY float64, renderer *sdl.Renderer) *SpriteButton {

	backgroundTex := texturedrawing.CreateSinglePixelTexture(backgroundColor, renderer)
	animTex := texturedrawing.CreateSinglePixelTexture(animBackgroundColor, renderer)
	selectedTex := texturedrawing.CreateSinglePixelTexture(selectedColor, renderer)

	sprite := sprite.NewSprite(path,
		pos,
		vec3.Vector3{X: 0, Y: 0, Z: 0},
		w,
		h,
		scaleX,
		scaleY,
		2,
		1,
		0,
		0,
		true,
		0,
		false,
		renderer)

	borderOffset := int32(float32(w) * borderPct)

	backPos := vec3.Vector3{X: float32(int32(pos.X) - borderOffset), Y: float32(int32(pos.Y) - borderOffset), Z: 0}
	spritePos := vec3.Vector3{X: pos.X, Y: pos.Y, Z: 0}
	width := (int32(w) + borderOffset*2)
	height := (int32(h) + borderOffset*2)
	rect := sdl.Rect{X: int32(backPos.X), Y: int32(backPos.Y), W: width, H: height}

	return &SpriteButton{sprite, rect, backgroundTex, animTex, selectedTex, false, false, false, animSpeedMS, false, 0, backPos, spritePos, width, height}
}

// Update updates whether the button was clicked or not
func (button *SpriteButton) Update(mouseState *MouseState, time float64) {
	if button.Rect.HasIntersection(&sdl.Rect{X: int32(mouseState.X), Y: int32(mouseState.Y), W: 1, H: 1}) {
		button.WasLeftClicked = !mouseState.PrevLeftButton && mouseState.LeftButton
		button.WasRightClicked = !mouseState.PrevRightButton && mouseState.RightButton
		button.IsSelected = true
	} else {
		button.WasLeftClicked = false
		button.WasRightClicked = false
		button.IsSelected = false
	}

	if button.WasLeftClicked == true {
		button.Animating = true
	}

	if button.Animating == true {
		button.AnimTimer += time
		if button.AnimTimer >= float64(button.AnimSpeed) {
			button.Animating = false
			button.AnimTimer = 0
		}
	}
	button.MainSprite.Update(time)
}

// Draw draws the button to the screen
func (button *SpriteButton) Draw(renderer *sdl.Renderer) {
	if button.IsSelected {
		borderRect := button.Rect
		borderThickness := int32(float32(borderRect.W) * 0.01)
		borderRect.W = button.Rect.W + borderThickness*2
		borderRect.H = button.Rect.H + borderThickness*2
		borderRect.X -= borderThickness
		borderRect.Y -= borderThickness
		renderer.Copy(button.SelectedTex, nil, &borderRect)
	}
	if button.Animating == true {
		button.MainSprite.CFrame = 1
		renderer.Copy(button.AnimBackground, nil, &button.Rect)
	} else {
		button.MainSprite.CFrame = 0
		renderer.Copy(button.Background, nil, &button.Rect)
	}
	button.MainSprite.Draw(renderer)
}

// TextButton structs contain all the information needed for a simple button with an image
type TextButton struct {
	WinWidth        int
	WinHeight       int
	Text            *font.TTFString
	Rect            sdl.Rect
	Background      *sdl.Texture
	AnimBackground  *sdl.Texture
	SelectedTex     *sdl.Texture
	IsSelected      bool
	WasLeftClicked  bool
	WasRightClicked bool
	AnimSpeed       int
	Animating       bool
	AnimTimer       float64
	BackgroundPos   vec3.Vector3
	TextPos         vec3.Vector3
	W               int
	H               int
	BorderOffset    int
}

// NewTextButton is a 'constructor' for a TextButton struct
func NewTextButton(winWidth, winHeight int, stringText string, size font.TextSize, textColor, backgroundColor, animBackgroundColor, selectedColor sdl.Color, pos vec3.Vector3, borderPct float32, animSpeedMS int, textFont *font.TTFFont, renderer *sdl.Renderer) *TextButton {

	backgroundTex := texturedrawing.CreateSinglePixelTexture(backgroundColor, renderer)
	animTex := texturedrawing.CreateSinglePixelTexture(animBackgroundColor, renderer)
	selectedTex := texturedrawing.CreateSinglePixelTexture(selectedColor, renderer)

	text := font.NewTTFString(stringText, size, textColor, pos, textFont, renderer)

	_, _, w, h, err := text.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	var borderOffset int
	switch size {
	case font.FontSmall:
		borderOffset = int(float32(textFont.SizeSmall) * borderPct)
	case font.FontMedium:
		borderOffset = int(float32(textFont.SizeMedium) * borderPct)
	case font.FontLarge:
		borderOffset = int(float32(textFont.SizeLarge) * borderPct)
	default:
		borderOffset = int(float32(textFont.SizeSmall) * borderPct)
	}

	backPos := vec3.Vector3{X: float32(pos.X - float32(borderOffset)), Y: float32(pos.Y - float32(borderOffset)), Z: 0}
	textPos := vec3.Vector3{X: pos.X, Y: pos.Y, Z: 0}
	width := int(w) + borderOffset*2
	height := int(h) + borderOffset*2
	rect := sdl.Rect{X: int32(backPos.X), Y: int32(backPos.Y), W: int32(width), H: int32(height)}

	return &TextButton{winWidth,
		winHeight,
		text,
		rect,
		backgroundTex,
		animTex,
		selectedTex,
		false,
		false,
		false,
		animSpeedMS,
		false,
		0,
		backPos,
		textPos,
		width,
		height,
		borderOffset}
}

// SetButtonPosition sets the positions of all the components of a button
func (button *TextButton) SetButtonPosition(pos vec3.Vector3) {
	button.TextPos = pos
	button.Text.Pos = button.TextPos
	button.BackgroundPos = vec3.Vector3{X: float32(button.TextPos.X - float32(button.BorderOffset)), Y: float32(button.TextPos.Y - float32(button.BorderOffset)), Z: 0}
	button.Rect = sdl.Rect{X: int32(button.BackgroundPos.X), Y: int32(button.BackgroundPos.Y), W: int32(button.W), H: int32(button.H)}
}

// SetCenterX sets the position to the center of the screen
func (button *TextButton) SetCenterX() {

	_, _, w, _, err := button.Text.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	if int(w) < button.WinWidth {
		diff := button.WinWidth - int(w)
		button.SetButtonPosition(vec3.Vector3{X: float32(diff / 2), Y: button.TextPos.Y, Z: 0})
	} else {
		diff := int(w) - button.WinWidth
		button.SetButtonPosition(vec3.Vector3{X: float32(diff / 2), Y: button.TextPos.Y, Z: 0})
	}
}

// SetCenterY sets the position to the center of the screen
func (button *TextButton) SetCenterY() {

	_, _, _, h, err := button.Text.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	if int(h) < button.WinHeight {
		diff := button.H - int(h)
		button.SetButtonPosition(vec3.Vector3{X: button.TextPos.X, Y: float32(diff / 2), Z: 0})
	} else {
		diff := int(h) - button.WinHeight
		button.SetButtonPosition(vec3.Vector3{X: button.TextPos.X, Y: float32(diff / 2), Z: 0})
	}
}

// Update updates whether the button was clicked or not
func (button *TextButton) Update(mouseState *MouseState, time float64) {
	if button.Rect.HasIntersection(&sdl.Rect{X: int32(mouseState.X), Y: int32(mouseState.Y), W: 1, H: 1}) {
		button.WasLeftClicked = !mouseState.PrevLeftButton && mouseState.LeftButton
		button.WasRightClicked = !mouseState.PrevRightButton && mouseState.RightButton
		button.IsSelected = true
	} else {
		button.WasLeftClicked = false
		button.WasRightClicked = false
		button.IsSelected = false
	}

	if button.WasLeftClicked == true {
		button.Animating = true
	}

	if button.Animating == true {
		button.AnimTimer += time
		if button.AnimTimer >= float64(button.AnimSpeed) {
			button.Animating = false
			button.AnimTimer = 0
		}
	}
}

// Draw draws the button to the screen
func (button *TextButton) Draw(renderer *sdl.Renderer) {
	if button.IsSelected {
		borderRect := button.Rect
		borderThickness := int32(float32(borderRect.W) * 0.01)
		borderRect.W = button.Rect.W + borderThickness*2
		borderRect.H = button.Rect.H + borderThickness*2
		borderRect.X -= borderThickness
		borderRect.Y -= borderThickness
		renderer.Copy(button.SelectedTex, nil, &borderRect)
	}
	if button.Animating == true {
		renderer.Copy(button.AnimBackground, nil, &button.Rect)
	} else {
		renderer.Copy(button.Background, nil, &button.Rect)
	}
	button.Text.Draw(renderer)
}
