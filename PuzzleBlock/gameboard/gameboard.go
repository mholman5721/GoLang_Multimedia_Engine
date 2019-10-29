package gameboard

import (
	"golang-games/PuzzleBlock/font"
	"golang-games/PuzzleBlock/gamestate"
	"golang-games/PuzzleBlock/gamestatetransition"
	"golang-games/PuzzleBlock/musicplayer"
	"golang-games/PuzzleBlock/soundplayer"
	"golang-games/PuzzleBlock/sprite"
	"math/rand"
	"strconv"

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
	CurrentGameState           *gamestatetransition.GameStateTransition
	MusicPlayer                *musicplayer.MusicPlayer
	SoundPlayer                *soundplayer.SoundPlayer
	Blocks                     [][]Block
	Background                 *sprite.Sprite
	LevelValue                 int
	MaxLevelValue              int
	PrevLevelValue             int
	ScoreValue                 int
	MaxScoreValue              int
	PrevScoreValue             int
	DeGrayValue                int
	MaxDeGrayValue             int
	PrevDeGrayValue            int
	BlockPointValue            int
	LevelScoreValue            int
	MaxLevelScoreValue         int
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

// Update updates all the tiles in the gameboard
func (g *GameBoard) Update(time float64) {

	// Update the background image
	g.Background.Update(time)

	// Move the current block down at a rate equal to the games current level
	if g.LevelFall == false && g.LevelFallingTimer >= g.LevelFallingTime {
		g.MoveActiveBlock("down")
		g.LevelFall = true
		g.LevelFallingTimer = 0
	} else if g.LevelFall == false && g.LevelFallingTimer < g.LevelFallingTime {
		g.LevelFallingTimer += time + (float64(g.LevelValue-1) * time)
	}

	// Make sure there is a 'time buffer' between the last time we pressed 'down' and the next time the active block automatically falls
	if g.LevelFall == true && g.LevelPostFallTimer >= g.LevelPostFallTime {
		g.LevelFall = false
		g.LevelPostFallTime = float64(g.MaxLevelValue*(g.MaxLevelValue-g.LevelValue)) + 1
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
				g.LevelValue = 1
				g.ScoreValue = 0
				g.DeGrayValue = 10
				g.LevelFallingTime = float64(g.MaxLevelValue * 100)
				g.LevelFallingTimer = 0
				g.LevelPostFallTimer = 0
				g.BlockFallingTimer = 0
				g.BlocksFallingTimer = 0

				// Change the game state
				g.MusicPlayer.FutureTune = 0
				g.CurrentGameState.TransitioningUp = true
				g.CurrentGameState.ToState = gamestate.TitleScreen
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

	// Update text - change levels if level points are above the number of points to change the level
	if g.LevelScoreValue >= g.MaxLevelScoreValue && g.LevelValue < g.MaxLevelValue {
		g.LevelScoreValue -= g.MaxLevelScoreValue
		g.LevelValue++
	}

	// Update text - DeGray the level if DeGray value is 0 or less
	if g.DeGrayValue <= 0 {
		for j := range g.BlockStates {
			for i := range g.BlockStates[j] {
				if g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence == 5 {
					g.Blocks[j][g.BlockStatesToGameBoard(i)].MainSprite.CSequence = rand.Intn(5)
					g.SetBlockColoring(g.BlockStatesToGameBoard(i), j)
				}
			}
		}
		g.DeGrayValue = g.MaxDeGrayValue
	}
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

	// Change the display text depending on whether the underlying value has changed
	if g.LevelValue != g.PrevLevelValue {
		g.LevelValueText.ChangeStringTexture(strconv.Itoa(g.LevelValue), font.FontLarge, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		g.PrevLevelValue = g.LevelValue
	}

	if g.ScoreValue != g.PrevScoreValue {
		g.ScoreValueText.ChangeStringTexture(strconv.Itoa(g.ScoreValue), font.FontLarge, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		g.PrevScoreValue = g.ScoreValue
	}

	if g.DeGrayValue != g.PrevDeGrayValue {
		g.DeGrayValueText.ChangeStringTexture(strconv.Itoa(g.DeGrayValue), font.FontLarge, sdl.Color{R: 255, G: 255, B: 255, A: 255}, renderer)
		g.PrevDeGrayValue = g.DeGrayValue
	}

	// Draw the text
	g.LevelText.Draw(renderer)
	g.LevelValueText.Draw(renderer)
	g.ScoreText.Draw(renderer)
	g.ScoreValueText.Draw(renderer)
	g.NextText.Draw(renderer)
	g.DeGrayText.Draw(renderer)
	g.DeGrayValueText.Draw(renderer)
}
