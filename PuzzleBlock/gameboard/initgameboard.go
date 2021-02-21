package gameboard

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/musicplayer"
	"golang-games/PuzzleBlock/soundplayer"
	"golang-games/PuzzleBlock/sprite"
	"golang-games/PuzzleBlock/vec3"
	"math/rand"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

// NewGameBoard is a gameboard constructor
func NewGameBoard(winWidth, winHeight, winDepth int, gamestate *gamestatetransition.GameStateTransition, numAcross, numDown, playAreaStart, playAreaEnd int, musicplayer *musicplayer.MusicPlayer, soundplayer *soundplayer.SoundPlayer, renderer *sdl.Renderer) *GameBoard {

	g := &GameBoard{}

	g.CurrentGameState = gamestate

	g.MusicPlayer = musicplayer

	g.SoundPlayer = soundplayer

	g.Blocks = make([][]Block, numDown)

	// Set blocks and explosions
	for j := 0; j < numDown; j++ {
		g.Blocks[j] = make([]Block, numAcross)
		for i := 0; i < numAcross; i++ {
			g.Blocks[j][i].MainSprite = sprite.NewSprite(
				"assets/Gems.png",
				vec3.Vector3{
					X: (float32(i) * float32(64)) * (float32(winWidth) / float32(numAcross)) / float32(64),
					Y: (float32(j) * float32(64)) * (float32(winHeight) / float32(numDown)) / float32(64),
					Z: float32(winDepth)},
				vec3.Vector3{
					X: 0,
					Y: 0,
					Z: 0},
				64,
				64,
				float64(winWidth/numAcross)/64,
				float64(winHeight/numDown)/64,
				10,
				7,
				0,
				6,
				true,
				100,
				true,
				renderer)
			g.Blocks[j][i].MainSprite.SetColorAndAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 255})

			g.Blocks[j][i].NumberOfExplosionFragments = rand.Intn(5) + 5
			g.Blocks[j][i].ExplosionSprites = make([]ExplosionSprite, g.Blocks[j][i].NumberOfExplosionFragments)
			for l := range g.Blocks[j][i].ExplosionSprites {

				g.Blocks[j][i].ExplosionSprites[l].OriginalPosition = FPos{
					X: ((float32(i) * float32(64)) * (float32(winWidth) / float32(numAcross)) / float32(64)) +
						(float32(32) * (float32(winWidth/numAcross) / 64)) +
						(float32(32-rand.Intn(64)-8) * (float32(winWidth/numAcross) / 64)),
					Y: (float32(j)*float32(64))*((float32(winHeight)/float32(numDown))/float32(64)) +
						(float32(32) * (float32(winHeight/numDown) / 64)) +
						(float32(32-rand.Intn(64)-8) * (float32(winWidth/numAcross) / 64))}

				g.Blocks[j][i].ExplosionSprites[l].MainSprite = sprite.NewSprite("assets/Gem.png",
					vec3.Vector3{
						X: float32(g.Blocks[j][i].ExplosionSprites[l].OriginalPosition.X),
						Y: float32(g.Blocks[j][i].ExplosionSprites[l].OriginalPosition.Y),
						Z: float32(winDepth)},
					vec3.Vector3{
						X: 0, //float32(rand.Intn(3)-1) / float32(rand.Intn(8)+1),
						Y: 0, //float32(rand.Intn(3)-1) / float32(rand.Intn(8)+1),
						Z: 0},
					16,
					16,
					float64(winWidth/numAcross)/64,
					float64(winHeight/numDown)/64,
					4,
					4,
					i%4,
					j%4,
					false,
					100,
					false,
					renderer)
				g.Blocks[j][i].ExplosionSprites[l].LifeSpan = float64(rand.Intn(100) + 150)
			}
		}
	}

	// Left hand side of gameboard
	for j := 0; j < numDown; j++ {
		for i := 0; i < playAreaStart; i++ {
			g.Blocks[j][i].MainSprite.CSequence = 5
			g.Blocks[j][i].MainSprite.Animating = false
			g.Blocks[j][i].MainSprite.CFrame = rand.Intn(10)
		}
	}

	// Center of gameboard
	for j := 0; j < numDown; j++ {
		for i := playAreaStart; i < playAreaEnd; i++ {

			g.Blocks[j][i].MainSprite.Drawing = false

			g.SetBlockColoring(i, j)
		}
	}

	// Right hand side of gameboard
	for j := 0; j < numDown; j++ {
		for i := playAreaEnd; i < numAcross; i++ {
			g.Blocks[j][i].MainSprite.CSequence = 5
			g.Blocks[j][i].MainSprite.Animating = false
			g.Blocks[j][i].MainSprite.CFrame = rand.Intn(10)
		}
	}

	// Score area
	for j := 1; j < 3; j++ {
		for i := 1; i < playAreaStart-1; i++ {
			g.Blocks[j][i].MainSprite.SetAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 64})
		}
	}

	// Level area
	for j := 4; j < 6; j++ {
		for i := 1; i < playAreaStart-1; i++ {
			g.Blocks[j][i].MainSprite.SetAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 64})
		}
	}

	// Next block area
	for j := 1; j < 3; j++ {
		for i := playAreaEnd + 1; i < numAcross-1; i++ {
			g.Blocks[j][i].MainSprite.SetAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 64})
		}
	}

	// De-gray area
	for j := 4; j < 6; j++ {
		for i := playAreaEnd + 1; i < numAcross-1; i++ {
			g.Blocks[j][i].MainSprite.SetAlpha(sdl.Color{R: 255, G: 255, B: 255, A: 64})
		}
	}

	// Set 'next' sprite
	g.Blocks[2][(numAcross+playAreaEnd)/2].MainSprite.CSequence = rand.Intn(7)
	g.SetBlockColoring((numAcross+playAreaEnd)/2, 2)
	g.Blocks[2][(numAcross+playAreaEnd)/2].MainSprite.Animating = true

	g.Background = sprite.NewSprite(
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

	g.LevelValue = 1
	g.MaxLevelValue = 10
	g.PrevLevelValue = g.LevelValue

	g.ScoreValue = 0
	g.MaxScoreValue = 9999999
	g.PrevScoreValue = g.ScoreValue

	g.DeGrayValue = 10
	g.MaxDeGrayValue = 10
	g.PrevDeGrayValue = g.DeGrayValue

	g.BlockPointValue = 10

	g.LevelScoreValue = 0
	g.MaxLevelScoreValue = 100

	g.NumAcross = numAcross
	g.NumDown = numDown
	g.PlayAreaStart = playAreaStart
	g.PlayAreaEnd = playAreaEnd

	g.ColorR = rand.Intn(256)
	g.ColorG = rand.Intn(256)
	g.ColorB = rand.Intn(256)
	g.ColorTimer = 0.0

	g.BlockStates = make([][]BlockState, numDown)
	for j := range g.BlockStates {
		g.BlockStates[j] = make([]BlockState, playAreaEnd-playAreaStart)
		for i := range g.BlockStates[j] {
			g.BlockStates[j][i] = Empty
		}
	}

	g.CurrentActive = Pos{-1, -1}

	// Set the font for the text
	g.TextFont = font.NewTTFFont("assets/FifteenTwenty-Bold.otf", winWidth, winHeight)

	// Set where the text goes on the screen
	g.ScoreText = font.NewTTFString("Score:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[1][1].MainSprite.Pos.X, Y: g.Blocks[1][1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.ScoreValueText = font.NewTTFString(strconv.Itoa(g.ScoreValue),
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[2][1].MainSprite.Pos.X, Y: g.Blocks[2][1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.LevelText = font.NewTTFString("Level:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[4][1].MainSprite.Pos.X, Y: g.Blocks[4][1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.LevelValueText = font.NewTTFString(strconv.Itoa(g.LevelValue),
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[5][3].MainSprite.Pos.X, Y: g.Blocks[5][3].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.NextText = font.NewTTFString("Next:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[1][playAreaEnd+1].MainSprite.Pos.X, Y: g.Blocks[1][playAreaEnd+1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.DeGrayText = font.NewTTFString("De-Gray:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[4][playAreaEnd+1].MainSprite.Pos.X, Y: g.Blocks[4][playAreaEnd+1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.DeGrayValueText = font.NewTTFString(strconv.Itoa(g.DeGrayValue),
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[5][playAreaEnd+3].MainSprite.Pos.X, Y: g.Blocks[5][playAreaEnd+3].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.LevelFallingTime = float64(g.MaxLevelValue * 100)
	g.LevelFallingTimer = 0

	g.LevelPostFallTime = float64(g.MaxLevelValue*(g.MaxLevelValue-g.LevelValue)) + 1
	g.LevelPostFallTimer = 0

	g.BlocksFalling = 0
	g.BlockFallingTime = 75
	g.BlockFallingTimer = 0

	g.BlocksFallingTime = g.BlockFallingTime * float64(g.NumDown)
	g.BlocksFallingTimer = 0

	g.GameOverTime = 1000
	g.GameOverTimer = 0

	g.GameOverPausing = false
	g.BlockScorePausing = false

	g.BlocksForScore = 0

	return g
}
