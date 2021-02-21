package font

import (
	"golang-games/PuzzleBlock/vec3"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// TextSize is an enum for the sizes of fonts
type TextSize int

const (
	// FontSmall is for small fonts
	FontSmall TextSize = iota
	// FontMedium is for medium fonts
	FontMedium
	// FontLarge is for large fonts
	FontLarge
	// FontTitle is for title fonts
	FontTitle
)

// TTFFont holds all the font information for TTFStrings
type TTFFont struct {
	FontSmall  *ttf.Font
	FontMedium *ttf.Font
	FontLarge  *ttf.Font
	FontTitle  *ttf.Font
	WinWidth   int
	WinHeight  int
	SizeSmall  int
	SizeMedium int
	SizeLarge  int
	SizeTitle  int
}

// TTFString is a struct that holds everything needed to draw text to the screen
type TTFString struct {
	Pos               vec3.Vector3
	Font              *TTFFont
	StringTexture     *sdl.Texture
	StringBackTexture *sdl.Texture
	StringText        string
}

// NewTTFFont Creates a new font object
func NewTTFFont(fontLocation string, winWidth, winHeight int) *TTFFont {

	font := &TTFFont{}

	var err error

	font.SizeSmall = int(float64(winWidth) * 0.015)

	font.FontSmall, err = ttf.OpenFont(fontLocation, font.SizeSmall)
	if err != nil {
		panic(err)
	}

	font.SizeMedium = int(float64(winWidth) * 0.03)

	font.FontMedium, err = ttf.OpenFont(fontLocation, font.SizeMedium)
	if err != nil {
		panic(err)
	}

	font.SizeLarge = int(float64(winWidth) * 0.06)

	font.FontLarge, err = ttf.OpenFont(fontLocation, font.SizeLarge)
	if err != nil {
		panic(err)
	}

	font.SizeTitle = int(float64(winWidth) * 0.15)

	font.FontTitle, err = ttf.OpenFont(fontLocation, font.SizeTitle)
	if err != nil {
		panic(err)
	}

	font.WinWidth = winWidth
	font.WinHeight = winHeight

	return font
}

// NewTTFString returns a pointer to a Font struct
func NewTTFString(stringText string, size TextSize, color sdl.Color, pos vec3.Vector3, font *TTFFont, renderer *sdl.Renderer) *TTFString {

	newString := &TTFString{Pos: pos, Font: font, StringTexture: nil, StringBackTexture: nil, StringText: stringText}

	newString.ChangeStringTexture(stringText, size, color, renderer)

	return newString
}

// ChangeStringTexture changes the texture associated with a TTFString entity
func (s *TTFString) ChangeStringTexture(stringText string, size TextSize, color sdl.Color, renderer *sdl.Renderer) {

	var fontSurface *sdl.Surface
	var backSurface *sdl.Surface
	var err error

	switch size {
	case FontSmall:
		fontSurface, err = s.Font.FontSmall.RenderUTF8Blended(stringText, color)
		if err != nil {
			panic(err)
		}
		backSurface, err = s.Font.FontSmall.RenderUTF8Blended(stringText, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}
	case FontMedium:
		fontSurface, err = s.Font.FontMedium.RenderUTF8Blended(stringText, color)
		if err != nil {
			panic(err)
		}
		backSurface, err = s.Font.FontMedium.RenderUTF8Blended(stringText, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}
	case FontLarge:
		fontSurface, err = s.Font.FontLarge.RenderUTF8Blended(stringText, color)
		if err != nil {
			panic(err)
		}
		backSurface, err = s.Font.FontLarge.RenderUTF8Blended(stringText, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}
	case FontTitle:
		fontSurface, err = s.Font.FontTitle.RenderUTF8Blended(stringText, color)
		if err != nil {
			panic(err)
		}
		backSurface, err = s.Font.FontTitle.RenderUTF8Blended(stringText, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}
	default:
		fontSurface, err = s.Font.FontSmall.RenderUTF8Blended(stringText, color)
		if err != nil {
			panic(err)
		}
		backSurface, err = s.Font.FontSmall.RenderUTF8Blended(stringText, sdl.Color{R: 0, G: 0, B: 0, A: 255})
		if err != nil {
			panic(err)
		}
	}

	tex, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		panic(err)
	}

	backTex, err := renderer.CreateTextureFromSurface(backSurface)
	if err != nil {
		panic(err)
	}

	s.StringTexture = tex
	s.StringBackTexture = backTex
}

// SetCenterX sets the position to the center of the screen
func (s *TTFString) SetCenterX() {

	_, _, w, _, err := s.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	if int(w) < s.Font.WinWidth {
		diff := s.Font.WinWidth - int(w)
		s.Pos.X = float32(diff / 2)
	} else {
		diff := int(w) - s.Font.WinWidth
		s.Pos.X = float32(diff / 2)
	}
}

// SetCenterY sets the position to the center of the screen
func (s *TTFString) SetCenterY() {

	_, _, _, h, err := s.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	if int(h) < s.Font.WinHeight {
		diff := s.Font.WinHeight - int(h)
		s.Pos.Y = float32(diff / 2)
	} else {
		diff := int(h) - s.Font.WinHeight
		s.Pos.Y = float32(diff / 2)
	}
}

// Draw draws the text to the screen
func (s *TTFString) Draw(renderer *sdl.Renderer) {

	_, _, w, h, err := s.StringTexture.Query()
	if err != nil {
		panic(err)
	}

	renderer.Copy(s.StringBackTexture, nil, &sdl.Rect{X: int32(s.Pos.X) + int32(float64(s.Font.WinWidth)*0.003), Y: int32(s.Pos.Y) + int32(float64(s.Font.WinWidth)*0.003), W: w, H: h})
	renderer.Copy(s.StringTexture, nil, &sdl.Rect{X: int32(s.Pos.X), Y: int32(s.Pos.Y), W: w, H: h})
}
