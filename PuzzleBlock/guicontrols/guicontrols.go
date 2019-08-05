package guicontrols

import (
	"PuzzleBlock/font"
	"PuzzleBlock/sprite"
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

// NewSpriteButton is a 'constructor' for an ImageButton struct
func NewSpriteButton(path string, backgroundColor, animBackgroundColor, selectedColor sdl.Color, pos vec3.Vector3, borderPct float32, animSpeedMS, w, h int, scaleX, scaleY float64, renderer *sdl.Renderer) *SpriteButton {
	backgroundTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton backgroundTex texture not created correctly...")
	}
	backgroundTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	backgroundPixels := make([]byte, 4)
	backgroundPixels[0] = backgroundColor.R
	backgroundPixels[1] = backgroundColor.G
	backgroundPixels[2] = backgroundColor.B
	backgroundPixels[3] = backgroundColor.A
	backgroundTex.Update(nil, backgroundPixels, 4)

	animTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton animTex texture not created correctly...")
	}
	animTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	animPixels := make([]byte, 4)
	animPixels[0] = animBackgroundColor.R
	animPixels[1] = animBackgroundColor.G
	animPixels[2] = animBackgroundColor.B
	animPixels[3] = animBackgroundColor.A
	animTex.Update(nil, animPixels, 4)

	selectedTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton selectedTex texture not created correctly...")
	}
	selectedTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	selectedPixels := make([]byte, 4)
	selectedPixels[0] = selectedColor.R
	selectedPixels[1] = selectedColor.G
	selectedPixels[2] = selectedColor.B
	selectedPixels[3] = selectedColor.A
	selectedTex.Update(nil, selectedPixels, 4)

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
		button.WasLeftClicked = mouseState.PrevLeftButton && !mouseState.LeftButton
		button.WasRightClicked = mouseState.PrevRightButton && !mouseState.RightButton
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
	W               int32
	H               int32
}

// NewTextButton is a 'constructor' for an ImageButton struct
func NewTextButton(stringText string, size font.TextSize, textColor, backgroundColor, animBackgroundColor, selectedColor sdl.Color, pos vec3.Vector3, borderPct float32, animSpeedMS int, textFont *font.TTFFont, renderer *sdl.Renderer) *TextButton {
	backgroundTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton backgroundTex texture not created correctly...")
	}
	backgroundTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	backgroundPixels := make([]byte, 4)
	backgroundPixels[0] = backgroundColor.R
	backgroundPixels[1] = backgroundColor.G
	backgroundPixels[2] = backgroundColor.B
	backgroundPixels[3] = backgroundColor.A
	backgroundTex.Update(nil, backgroundPixels, 4)

	animTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton animTex texture not created correctly...")
	}
	animTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	animPixels := make([]byte, 4)
	animPixels[0] = animBackgroundColor.R
	animPixels[1] = animBackgroundColor.G
	animPixels[2] = animBackgroundColor.B
	animPixels[3] = animBackgroundColor.A
	animTex.Update(nil, animPixels, 4)

	selectedTex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic("ERROR: NewTextButton selectedTex texture not created correctly...")
	}
	selectedTex.SetBlendMode(sdl.BLENDMODE_BLEND)

	selectedPixels := make([]byte, 4)
	selectedPixels[0] = selectedColor.R
	selectedPixels[1] = selectedColor.G
	selectedPixels[2] = selectedColor.B
	selectedPixels[3] = selectedColor.A
	selectedTex.Update(nil, selectedPixels, 4)

	text := font.NewTTFString(stringText, size, textColor, pos, textFont, renderer)

	_, _, w, h, err := text.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	var borderOffset int32
	switch size {
	case font.FontSmall:
		borderOffset = int32(float32(textFont.SizeSmall) * borderPct)
	case font.FontMedium:
		borderOffset = int32(float32(textFont.SizeMedium) * borderPct)
	case font.FontLarge:
		borderOffset = int32(float32(textFont.SizeLarge) * borderPct)
	default:
		borderOffset = int32(float32(textFont.SizeSmall) * borderPct)
	}

	backPos := vec3.Vector3{X: float32(int32(pos.X) - borderOffset), Y: float32(int32(pos.Y) - borderOffset), Z: 0}
	textPos := vec3.Vector3{X: pos.X, Y: pos.Y, Z: 0}
	width := (w + borderOffset*2)
	height := (h + borderOffset*2)
	rect := sdl.Rect{X: int32(backPos.X), Y: int32(backPos.Y), W: width, H: height}

	return &TextButton{text, rect, backgroundTex, animTex, selectedTex, false, false, false, animSpeedMS, false, 0, backPos, textPos, width, height}
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

// GetSinglePixelTex returns a texture consisting of a single colored pixel
func GetSinglePixelTex(renderer *sdl.Renderer, color sdl.Color) *sdl.Texture {
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic(err)
	}

	pixels := make([]byte, 4)
	pixels[0] = color.R
	pixels[1] = color.G
	pixels[2] = color.B
	pixels[3] = color.A
	tex.Update(nil, pixels, 4)

	return tex
}
