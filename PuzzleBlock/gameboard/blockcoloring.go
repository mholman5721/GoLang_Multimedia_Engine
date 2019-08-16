package gameboard

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

// SetBlockColoring sets the color and alpha values of a single block
func (g *GameBoard) SetBlockColoring(i, j int) {

	switch g.Blocks[j][i].MainSprite.CSequence {
	case 0: // RED
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 0, B: 0, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 0, B: 0, A: 128})
		}
	case 1: // GREEN
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 0, G: 255, B: 0, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 0, G: 255, B: 0, A: 128})
		}
	case 2: // BLUE
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 0, G: 0, B: 255, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 0, G: 0, B: 255, A: 128})
		}
	case 3: // YELLOW
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 0, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 0, A: 128})
		}
	case 4: // VIOLET
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 128, G: 0, B: 128, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 128, G: 0, B: 128, A: 128})
		}
	case 5: // GRAY
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 128, G: 128, B: 128, A: 255})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 128, G: 128, B: 128, A: 128})
		}
	case 6: // MULTI
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 128})
		}
	default:
		g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 128})
		for k := range g.Blocks[j][i].ExplosionSprites {
			g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 128})
		}
	}
}

// UpdateMultiBlockColor is a helper function which updates the colors of the multiblocks
func (g *GameBoard) UpdateMultiBlockColor(i, j int) {
	g.ColorR += rand.Intn(64)
	g.ColorG += rand.Intn(64)
	g.ColorB += rand.Intn(64)
	g.ColorR %= 256
	g.ColorG %= 256
	g.ColorB %= 256
	g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: uint8(g.ColorR), G: uint8(g.ColorG), B: uint8(g.ColorB), A: 128})
}
