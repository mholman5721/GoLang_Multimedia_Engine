package texturedrawing

import "github.com/veandco/go-sdl2/sdl"

// SinglePixelTexture contains the data for making monocolor rectangular textures
type SinglePixelTexture struct {
	Rect    sdl.Rect
	Texture *sdl.Texture
}

// NewSinglePixelTexture returns a texture consisting of a single colored pixel
func NewSinglePixelTexture(color sdl.Color, rect sdl.Rect, renderer *sdl.Renderer) *SinglePixelTexture {

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	if err != nil {
		panic(err)
	}
	tex.SetBlendMode(sdl.BLENDMODE_BLEND)

	data := make([]byte, 4)
	data[3] = color.A
	data[2] = color.B
	data[1] = color.G
	data[0] = color.R

	tex.Update(nil, data, 4)

	t := &SinglePixelTexture{rect, tex}

	return t
}

// Draw draws the texture
func (t *SinglePixelTexture) Draw(renderer *sdl.Renderer) {
	renderer.Copy(t.Texture, nil, &t.Rect)
}
