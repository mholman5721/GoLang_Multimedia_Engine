package sprite

import (
	"vec3"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Sprite is a struct that contains the basic building blocks for game entities
type Sprite struct {
	Tex                 *sdl.Texture
	Src                 *sdl.Rect
	Dst                 *sdl.Rect
	Pos                 vec3.Vector3
	Vel                 vec3.Vector3
	W, H                int
	ScaleX, ScaleY      float64
	NFrames, NSequences int
	CFrame, CSequence   int
	Drawing             bool
	AnimSpeed           int
	Animating           bool
	AnimTimer           float64
}

// NewSprite returns a pointer to a newly created Sprite object
func NewSprite(path string, pos, vel vec3.Vector3, w, h int, scaleX, scaleY float64, nFrames, nSequences, cFrame, cSequence int, drawing bool, animSpeed int, animating bool, renderer *sdl.Renderer) *Sprite {

	s := &Sprite{}

	image, err := img.Load(path)
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		panic(err)
	}

	s.Tex = texture
	s.Pos = pos
	s.Vel = vel
	s.W = w
	s.H = h
	s.ScaleX = scaleX
	s.ScaleY = scaleY
	s.NFrames = nFrames
	s.NSequences = nSequences
	s.CFrame = cFrame
	s.CSequence = cSequence
	s.Drawing = drawing
	s.AnimSpeed = animSpeed
	s.Animating = animating
	s.AnimTimer = 0

	s.Src = &sdl.Rect{X: 0, Y: 0, W: int32(s.W), H: int32(s.H)}
	s.Dst = &sdl.Rect{X: int32(s.Pos.X), Y: int32(s.Pos.Y), W: int32(s.W), H: int32(s.H)}

	return s
}

// Update sets the sprite's src and dst rectangles depending on where it is in the animation
func (s *Sprite) Update(time float64) {
	if s.Animating == true && s.AnimSpeed > 0 {
		s.AnimTimer += time
		if s.AnimTimer >= float64(s.AnimSpeed) {
			s.CFrame++
			s.CFrame %= s.NFrames

			s.AnimTimer = 0
		}
	}

	s.Src.X = int32(s.CFrame * s.W)
	s.Src.Y = int32(s.CSequence * s.H)
	s.Src.W = int32(s.W)
	s.Src.H = int32(s.H)

	s.Pos.X += s.Vel.X
	s.Pos.Y += s.Vel.Y
	s.Pos.Z += s.Vel.Z

	s.Dst.X = int32(s.Pos.X)
	s.Dst.Y = int32(s.Pos.Y)
	s.Dst.W = int32(float64(s.W) * s.ScaleX)
	s.Dst.H = int32(float64(s.H) * s.ScaleY)
}

// SetColor sets the color values of a sprite
func (s *Sprite) SetColor(c sdl.Color) {
	s.Tex.SetColorMod(c.R, c.G, c.B)
}

// SetAlpha sets the alpha value of a sprite
func (s *Sprite) SetAlpha(c sdl.Color) {
	s.Tex.SetAlphaMod(c.A)
}

// SetColorAndAlpha sets the color and alpha values of a sprite
func (s *Sprite) SetColorAndAlpha(c sdl.Color) {
	s.Tex.SetColorMod(c.R, c.G, c.B)
	s.Tex.SetAlphaMod(c.A)
}

// Draw instructs the renderer to copy the sprite to the renderer buffer
func (s *Sprite) Draw(renderer *sdl.Renderer) {
	if s.Drawing == true {
		renderer.Copy(s.Tex, s.Src, s.Dst)
	}
}
