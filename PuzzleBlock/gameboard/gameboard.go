package gameboard

import (
	"PuzzleBlock/font"
	"PuzzleBlock/sprite"
	"math/rand"
	"vec3"

	"github.com/veandco/go-sdl2/sdl"
)

// Pos is a struct that holds X/ Y positions for the gameboard
type Pos struct {
	X, Y int
}

// FPos is a struct that holds X/ Y positions in floating point values
type FPos struct {
	X, Y float32
}

// BlockState denotes the states a block on the gameboard can be in
type BlockState int

const (
	// Empty block is a block that is not being drawn
	Empty BlockState = iota
	// Active block is controlled by the player
	Active
	// Inactive is a block that is being drawn, but which is not controlled by the player
	Inactive
	// Exploding is a block that is in the process of exploding
	Exploding
)

// ExplosionSprite contains the data needed for each fragment of an exploding block
type ExplosionSprite struct {
	MainSprite       *sprite.Sprite
	LifeSpan         float64
	CurrentLife      float64
	OriginalPosition FPos
}

// Block is the basic structure from which everything is made
type Block struct {
	MainSprite                 *sprite.Sprite
	ExplosionSprites           []ExplosionSprite
	NumberOfExplosionFragments int
}

// GameBoard is a struct that contains all the sprite information for the game
type GameBoard struct {
	Blocks                     [][]Block
	Background                 *sprite.Sprite
	LevelValue                 int
	MaxLevelValue              int
	NumAcross, NumDown         int
	PlayAreaStart, PlayAreaEnd int
	ColorR                     int
	ColorG                     int
	ColorB                     int
	ColorTimer                 float64
	BlockStates                [][]BlockState
	CurrentActive              Pos
	TextFont                   *font.TTFFont
	LevelText                  *font.TTFString
	LevelValueText             *font.TTFString
	ScoreText                  *font.TTFString
	ScoreValueText             *font.TTFString
	NextText                   *font.TTFString
	DeGrayText                 *font.TTFString
	DeGrayValueText            *font.TTFString
	TimeSinceLastDown          float64
	LevelFall                  bool
	LevelFallingTime           float64
	LevelFallingTimer          float64
	LevelPostFallTime          float64
	LevelPostFallTimer         float64
	BlocksFalling              int
	BlockFallingTime           float64
	BlockFallingTimer          float64
	BlocksFallingTime          float64
	BlocksFallingTimer         float64
	GameOverTime               float64
	GameOverTimer              float64
	GameOverPausing            bool
	BlockScorePausing          bool
}

// GameBoardToBlockStates translates an x coordinate in the play area to an x coordinate in the block states slice
func (g *GameBoard) GameBoardToBlockStates(i int) int {
	return i - g.PlayAreaStart - 1
}

// BlockStatesToGameBoard translates an x coordinate in the block states to an x coordinate in the play area slice
func (g *GameBoard) BlockStatesToGameBoard(i int) int {
	return i + g.PlayAreaStart
}

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

// NewGameBoard is a gameboard constructor
func NewGameBoard(winWidth, winHeight, winDepth, numAcross, numDown, playAreaStart, playAreaEnd int, renderer *sdl.Renderer) *GameBoard {

	g := &GameBoard{}

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
	g.TextFont = font.NewTTFFont("assets/FifteenTwenty-Bold.otf", winWidth)

	// Set where the text goes on the screen
	g.ScoreText = font.NewTTFString("Score:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[1][1].MainSprite.Pos.X, Y: g.Blocks[1][1].MainSprite.Pos.Y, Z: 0},
		g.TextFont,
		renderer)

	g.LevelText = font.NewTTFString("Level:",
		font.FontLarge,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
		vec3.Vector3{X: g.Blocks[4][1].MainSprite.Pos.X, Y: g.Blocks[4][1].MainSprite.Pos.Y, Z: 0},
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

	g.TimeSinceLastDown = 0

	g.LevelFallingTime = float64(g.MaxLevelValue * 100)
	g.LevelFallingTimer = 0

	g.LevelPostFallTime = 150
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

	return g
}

// ProccessBlockMovement handles the logic of moving the active block
func (g *GameBoard) ProccessBlockMovement(d string) {

	prevActive := g.CurrentActive

	switch d {
	case "Y++":
		if g.CurrentActive.Y >= 0 && g.CurrentActive.Y < g.NumDown && g.CurrentActive.Y+1 < g.NumDown &&
			g.BlockStates[g.CurrentActive.Y+1][g.CurrentActive.X] == Empty {
			g.CurrentActive.Y++
		}
	case "Y--":
		if g.CurrentActive.Y > 0 && g.CurrentActive.Y <= g.NumDown {
			g.CurrentActive.Y--
		}
	case "X++":
		if (g.CurrentActive.X >= 0 && g.CurrentActive.X < g.GameBoardToBlockStates(g.PlayAreaEnd)) &&
			g.BlockStates[g.CurrentActive.Y][g.CurrentActive.X+1] == Empty {
			g.CurrentActive.X++
		}
	case "X--":
		if (g.CurrentActive.X > 0 && g.CurrentActive.X <= g.GameBoardToBlockStates(g.PlayAreaEnd)) &&
			g.BlockStates[g.CurrentActive.Y][g.CurrentActive.X-1] == Empty {
			g.CurrentActive.X--
		}
	default:
		panic("ERROR: ProccessBlockMovement requires input of Y++, Y--, X++, X--. You have: " + d)
	}
	//fmt.Println(g.CurrentActive.X, g.CurrentActive.Y, g.GameBoardToBlockStates(g.PlayAreaEnd))

	if g.CurrentActive.Y < g.NumDown && g.CurrentActive.Y >= 0 {
		// Set the old block to empty and not drawing
		g.BlockStates[prevActive.Y][prevActive.X] = Empty
		g.Blocks[prevActive.Y][g.BlockStatesToGameBoard(prevActive.X)].MainSprite.Drawing = false

		// Set the current block to active and drawing
		g.BlockStates[g.CurrentActive.Y][g.CurrentActive.X] = Active
		g.Blocks[g.CurrentActive.Y][g.BlockStatesToGameBoard(g.CurrentActive.X)].MainSprite.CSequence = g.Blocks[prevActive.Y][g.BlockStatesToGameBoard(prevActive.X)].MainSprite.CSequence
		g.SetBlockColoring(g.BlockStatesToGameBoard(g.CurrentActive.X), g.CurrentActive.Y)
		g.Blocks[g.CurrentActive.Y][g.BlockStatesToGameBoard(g.CurrentActive.X)].MainSprite.Drawing = true
		//fmt.Println(g.ColorR, g.ColorG, g.ColorB)
	}
}

// MoveActiveBlock changes around the game map based on the user pressed key
func (g *GameBoard) MoveActiveBlock(d string) {
	if g.GameOverPausing == false {
		switch d {
		case "up":
			//fmt.Println("up")
			g.ProccessBlockMovement("Y--")
		case "down":
			//fmt.Println("down")
			g.ProccessBlockMovement("Y++")
			g.TimeSinceLastDown = 0
		case "left":
			//fmt.Println("left")
			g.ProccessBlockMovement("X--")
		case "right":
			//fmt.Println("right")
			g.ProccessBlockMovement("X++")
		default:
			//fmt.Println("else")
		}
	}
	/*
		for _, line := range g.BlockStates {
			fmt.Println(line)
		}
		fmt.Println()
	*/
}

// CheckScore checks the gameboard for scores in rows columns and diagonals
func (g *GameBoard) CheckScore(direction string, originalBlock, nextBlock Pos, score int) int {

	switch direction {
	case "up":
		// Not needed
	case "down":
		if nextBlock.Y > (g.NumDown - 1) {
			return score
		}
	case "left":
		if nextBlock.X < 0 {
			return score
		}
	case "right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 {
			return score
		}
	case "up_left":
		if nextBlock.X < 0 || nextBlock.Y < 0 {
			return score
		}
	case "up_right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 || nextBlock.Y < 0 {
			return score
		}
	case "down_left":
		if nextBlock.X < 0 || nextBlock.Y > (g.NumDown-1) {
			return score
		}
	case "down_right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 || nextBlock.Y > (g.NumDown-1) {
			return score
		}
	default:
	}

	if g.Blocks[originalBlock.Y][g.BlockStatesToGameBoard(originalBlock.X)].MainSprite.CSequence == g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence &&
		g.Blocks[originalBlock.Y][g.BlockStatesToGameBoard(originalBlock.X)].MainSprite.CSequence != 6 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence != 6 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence != 5 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.Drawing == true &&
		g.BlockStates[nextBlock.Y][nextBlock.X] == Inactive {

		score++
		g.BlockStates[originalBlock.Y][originalBlock.X] = Exploding
		g.BlockStates[nextBlock.Y][nextBlock.X] = Exploding

		switch direction {
		case "up":
			// Not needed
		case "down":
			score += g.CheckScore("down", originalBlock, Pos{nextBlock.X, nextBlock.Y + 1}, score)
		case "left":
			score += g.CheckScore("left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y}, score)
		case "right":
			score += g.CheckScore("right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y}, score)
		case "up_left":
			score += g.CheckScore("up_left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y - 1}, score)
		case "up_right":
			score += g.CheckScore("up_right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y - 1}, score)
		case "down_left":
			score += g.CheckScore("down_left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y + 1}, score)
		case "down_right":
			score += g.CheckScore("down_right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y + 1}, score)
		default:
		}
	}

	return score
}

// HandleScoreBlocks contains the logic for what should happen to blocks after they are marked by the CheckScore functions
func (g *GameBoard) HandleScoreBlocks(score int) {
	if score >= 3 {
		g.BlockScorePausing = true
		//fmt.Println("~~~~~SCORE: ", score)
		for k := range g.BlockStates {
			for l := range g.BlockStates[k] {
				if g.BlockStates[k][l] == Exploding {
					g.BlockStates[k][l] = Empty
					g.Blocks[k][g.BlockStatesToGameBoard(l)].MainSprite.Drawing = false
					for o := range g.Blocks[k][g.BlockStatesToGameBoard(l)].ExplosionSprites {
						g.Blocks[k][g.BlockStatesToGameBoard(l)].ExplosionSprites[o].MainSprite.Vel.X = float32(rand.Intn(3)-1) / float32(rand.Intn(8)+1)
						g.Blocks[k][g.BlockStatesToGameBoard(l)].ExplosionSprites[o].MainSprite.Vel.Y = float32(rand.Intn(3)-1) / float32(rand.Intn(8)+1)
						g.Blocks[k][g.BlockStatesToGameBoard(l)].ExplosionSprites[o].MainSprite.Drawing = true
					}
				}
			}
		}
	} else {
		for n := range g.BlockStates {
			for m := range g.BlockStates[n] {
				if g.BlockStates[n][m] == Exploding {
					g.BlockStates[n][m] = Inactive
				}
			}
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

// Update updates all the tiles in the gameboard
func (g *GameBoard) Update(time float64) {

	// Update the background image
	g.Background.Update(time)

	// Update the time since the last time the down key was pressed
	if g.TimeSinceLastDown >= g.LevelFallingTime {
		g.TimeSinceLastDown = 0
	} else {
		g.TimeSinceLastDown += time
	}

	// Move the current block down at a rate equal to the games current level
	if g.LevelFall == false && (g.LevelFallingTimer-g.TimeSinceLastDown) >= g.LevelFallingTime {
		g.MoveActiveBlock("down")
		g.LevelFall = true
		g.LevelFallingTimer = 0
	} else if g.LevelFall == false && (g.LevelFallingTimer-g.TimeSinceLastDown) < g.LevelFallingTime {
		if g.MaxLevelValue-g.LevelValue == 0 {
			g.LevelFallingTimer += 2 * time
		} else {
			g.LevelFallingTimer += time + (time / float64(g.MaxLevelValue-g.LevelValue))
		}
	}

	// Make sure there is a 'time buffer' between the last time we pressed 'down' and the next time the active block automatically falls
	if g.LevelFall == true && g.LevelPostFallTimer >= g.LevelPostFallTime {
		g.LevelFall = false
		g.LevelPostFallTimer = 0
	} else if g.LevelFall == true && g.LevelPostFallTimer < g.LevelPostFallTime {
		g.LevelPostFallTimer += time
	}

	// Stop the downward descent of the current block
	if g.CurrentActive.Y == g.NumDown-1 || ((g.CurrentActive.X != -1 && g.CurrentActive.Y != -1) && g.BlockStates[g.CurrentActive.Y+1][g.CurrentActive.X] == Inactive) {
		g.BlockStates[g.CurrentActive.Y][g.CurrentActive.X] = Inactive
		g.CurrentActive = Pos{-1, -1}
	}

	// Check for game over state which occurs when one column of blocks reaches the top of the gameboard
	currentYCount := make([]int, g.PlayAreaEnd-g.PlayAreaStart)
	for j := range g.BlockStates {
		for i := range g.BlockStates[j] {
			if g.BlockStates[j][i] == Inactive {
				currentYCount[i]++
			}
		}
	}

	// If a column of blocks reaches the top of the gameboard, reset everything for now
	// TODO: add 'game over' state/ screen and transition to that instead of resetting everything
	for k := range currentYCount {
		if currentYCount[k] >= g.NumDown {
			g.GameOverPausing = true
			if g.GameOverTimer >= g.GameOverTime {
				for j := range g.BlockStates {
					for i := range g.BlockStates[j] {
						g.BlockStates[j][i] = Empty
						g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing = false
					}
				}
				g.CurrentActive = Pos{-1, -1}
				g.GameOverTimer = 0
				g.GameOverPausing = false
				break
			} else {
				g.GameOverTimer += time
			}
		}
	}
	currentYCount = nil

	// Update the colors of the multi-blocks
	g.ColorTimer += time
	if g.ColorTimer >= 50 {
		for j := 0; j < g.NumDown; j++ {
			for i := g.PlayAreaStart; i < g.PlayAreaEnd; i++ {
				if g.Blocks[j][i].MainSprite.Drawing == true && g.Blocks[j][i].MainSprite.CSequence == 6 {
					g.UpdateMultiBlockColor(i, j)
				}
			}
		}
		if g.Blocks[2][(g.NumAcross+g.PlayAreaEnd)/2].MainSprite.CSequence == 6 {
			g.UpdateMultiBlockColor((g.NumAcross+g.PlayAreaEnd)/2, 2)
		}
		g.ColorTimer = 0
	}

	// Check for block scores
	score := 0
	for j := range g.BlockStates {
		for i := range g.BlockStates[j] {
			if g.BlockStates[j][i] == Inactive && g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == true {
				if g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence == 6 {
					if j+1 > g.NumDown-1 {
						for g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence == 6 {
							g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence = rand.Intn(7)
							g.SetBlockColoring(g.BlockStatesToGameBoard(i), j)
						}
					} else {
						g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence = g.Blocks[j+1][g.BlockStatesToGameBoard(i)].MainSprite.CSequence
						g.SetBlockColoring(g.BlockStatesToGameBoard(i), j)
					}

				}

				score += g.CheckScore("down", Pos{i, j}, Pos{i, j + 1}, score)
				//fmt.Println("Down SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("left", Pos{i, j}, Pos{i - 1, j}, score)
				//fmt.Println("Left SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("right", Pos{i, j}, Pos{i + 1, j}, score)
				//fmt.Println("Right SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("up_left", Pos{i, j}, Pos{i - 1, j - 1}, score)
				//fmt.Println("Right SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("up_right", Pos{i, j}, Pos{i + 1, j - 1}, score)
				//fmt.Println("Right SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("down_left", Pos{i, j}, Pos{i - 1, j + 1}, score)
				//fmt.Println("Right SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0

				score += g.CheckScore("down_right", Pos{i, j}, Pos{i + 1, j + 1}, score)
				//fmt.Println("Right SCORE: ", score)
				g.HandleScoreBlocks(score)
				score = 0
			}
		}
	}

	// Update explosion sprite lifespans and reset once they are done
	for j := range g.Blocks {
		for i := range g.Blocks[j] {
			for k := range g.Blocks[j][i].ExplosionSprites {
				if g.Blocks[j][i].ExplosionSprites[k].MainSprite.Drawing == true && g.Blocks[j][i].ExplosionSprites[k].CurrentLife < g.Blocks[j][i].ExplosionSprites[k].LifeSpan {
					g.Blocks[j][i].ExplosionSprites[k].CurrentLife += time
					alpha := uint8(((g.Blocks[j][i].ExplosionSprites[k].LifeSpan - g.Blocks[j][i].ExplosionSprites[k].CurrentLife) / g.Blocks[j][i].ExplosionSprites[k].LifeSpan) * 128)
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetAlpha(sdl.Color{R: 0, G: 0, B: 0, A: alpha})
				} else {
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.Pos.X = g.Blocks[j][i].ExplosionSprites[k].OriginalPosition.X
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.Pos.Y = g.Blocks[j][i].ExplosionSprites[k].OriginalPosition.Y
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.Vel.X = 0
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.Vel.Y = 0
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.Drawing = false
					g.Blocks[j][i].ExplosionSprites[k].CurrentLife = 0
					g.Blocks[j][i].ExplosionSprites[k].MainSprite.SetAlpha(sdl.Color{R: 0, G: 0, B: 0, A: 0})
				}
			}
		}
	}

	// Check for falling blocks
	if g.BlocksFalling == 0 {
		for j := range g.BlockStates {
			for i := range g.BlockStates[j] {
				if g.BlockStates[j][i] == Inactive &&
					g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == true &&
					(j+1 < g.NumDown) && (g.BlockStates[j+1][i] == Empty) &&
					(g.Blocks[j+1][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == false) {

					g.BlocksFalling++
					//g.BlockFallPausing = true
				}
			}
		}
	}

	// Update falling blocks and the ones below them
	if g.BlocksFalling > 0 && g.BlockFallingTimer >= g.BlockFallingTime {
		for j := range g.BlockStates {
			for i := range g.BlockStates[j] {
				if g.BlockStates[j][i] == Inactive && g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == true {
					if (j+1 < g.NumDown) && (g.BlockStates[j+1][i] == Empty) && (g.Blocks[j+1][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == false) {
						g.BlockStates[j][i] = Empty
						g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing = false

						g.BlockStates[j+1][i] = Inactive
						g.Blocks[j+1][g.BlockStatesToGameBoard(i)].MainSprite.CSequence = g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence
						g.SetBlockColoring(g.BlockStatesToGameBoard(i), j+1)
						g.Blocks[j+1][g.BlockStatesToGameBoard(i)].MainSprite.Drawing = true

						g.BlocksFalling--
						g.BlockFallingTimer = 0
					}
				}
			}
		}
	} else if g.BlocksFalling > 0 && g.BlockFallingTimer < g.BlockFallingTime {
		g.BlockFallingTimer += time
	}

	g.BlocksFalling = 0

	// Update the block falling timer
	if g.BlockScorePausing == true && g.BlocksFallingTimer >= g.BlocksFallingTime {
		g.BlockScorePausing = false
		g.BlocksFallingTimer = 0
	} else if g.BlockScorePausing == true && g.BlocksFallingTimer < g.BlocksFallingTime {
		g.BlocksFallingTimer += time
	}

	// Make sure empty blocks are empty
	for j := range g.BlockStates {
		for i := range g.BlockStates[j] {
			if g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.Drawing == false {
				g.BlockStates[j][i] = Empty
			}
		}
	}

	// Spawn a new current block at the top of the play area only once all other checks are complete
	if g.BlocksFalling == 0 &&
		g.BlockScorePausing == false &&
		(g.CurrentActive.X == -1 && g.CurrentActive.Y == -1) &&
		g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.Drawing == false {

		g.CurrentActive = Pos{2, 0}
		g.BlockStates[0][2] = Active
		g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.CSequence = g.Blocks[2][(g.NumAcross+g.PlayAreaEnd)/2].MainSprite.CSequence
		g.SetBlockColoring((g.PlayAreaStart+g.PlayAreaEnd)/2, 0)
		g.Blocks[2][(g.NumAcross+g.PlayAreaEnd)/2].MainSprite.CSequence = rand.Intn(7)
		g.SetBlockColoring((g.NumAcross+g.PlayAreaEnd)/2, 2)

		// Check if the block below the starting block is being drawn - ensure game over if it is
		if g.Blocks[1][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.Drawing == true {
			for g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.CSequence == 6 || g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.CSequence == g.Blocks[1][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.CSequence {
				g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.CSequence = rand.Intn(6)
			}
		}

		g.SetBlockColoring((g.PlayAreaStart+g.PlayAreaEnd)/2, 0)
		g.Blocks[0][(g.PlayAreaStart+g.PlayAreaEnd)/2].MainSprite.Drawing = true
	}

	// Update all the blocks
	for j := range g.Blocks {
		for i := range g.Blocks[j] {
			g.Blocks[j][i].MainSprite.Update(time)
		}
	}

	// Update explosion fragments
	for j := range g.Blocks {
		for i := range g.Blocks[j] {
			for k := range g.Blocks[j][i].ExplosionSprites {
				g.Blocks[j][i].ExplosionSprites[k].MainSprite.Update(time)
			}
		}
	}

	// Update text

}

// Draw draws all the tiles in the gameboard
func (g *GameBoard) Draw(renderer *sdl.Renderer) {

	// Draw the background
	g.Background.Draw(renderer)

	// Draw the blocks
	for j := range g.Blocks {
		for i := range g.Blocks[j] {
			g.Blocks[j][i].MainSprite.Draw(renderer)
		}
	}

	// Draw the explosion sprites
	for j := range g.Blocks {
		for i := range g.Blocks[j] {
			for k := range g.Blocks[j][i].ExplosionSprites {
				g.Blocks[j][i].ExplosionSprites[k].MainSprite.Draw(renderer)
			}
		}
	}

	// Draw the text
	g.LevelText.Draw(renderer)
	g.ScoreText.Draw(renderer)
	g.NextText.Draw(renderer)
	g.DeGrayText.Draw(renderer)
}
