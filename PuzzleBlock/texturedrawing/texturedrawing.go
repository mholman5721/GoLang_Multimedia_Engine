package texturedrawing

import "github.com/veandco/go-sdl2/sdl"

// CreateSinglePixelTexture returns a texture consisting of a single colored pixel
func CreateSinglePixelTexture(color sdl.Color, renderer *sdl.Renderer) *sdl.Texture {
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

	return tex
}
