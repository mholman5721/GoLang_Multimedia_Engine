package gameboard

import (
	"math/rand"
	"strconv"
)

// CheckScore checks the gameboard for scores in rows columns and diagonals
func (g *GameBoard) CheckScore(direction string, originalBlock, nextBlock Pos) {

	switch direction {
	case "up":
		// Not needed
	case "down":
		if nextBlock.Y > (g.NumDown - 1) {
			return
		}
	case "left":
		if nextBlock.X < 0 {
			return
		}
	case "right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 {
			return
		}
	case "up_left":
		if nextBlock.X < 0 || nextBlock.Y < 0 {
			return
		}
	case "up_right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 || nextBlock.Y < 0 {
			return
		}
	case "down_left":
		if nextBlock.X < 0 || nextBlock.Y > (g.NumDown-1) {
			return
		}
	case "down_right":
		if nextBlock.X > (g.PlayAreaEnd-g.PlayAreaStart)-1 || nextBlock.Y > (g.NumDown-1) {
			return
		}
	default:
	}

	if g.Blocks[originalBlock.Y][g.BlockStatesToGameBoard(originalBlock.X)].MainSprite.CSequence == g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence &&
		g.Blocks[originalBlock.Y][g.BlockStatesToGameBoard(originalBlock.X)].MainSprite.CSequence != 6 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence != 6 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(originalBlock.X)].MainSprite.CSequence != 5 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.CSequence != 5 &&
		g.Blocks[nextBlock.Y][g.BlockStatesToGameBoard(nextBlock.X)].MainSprite.Drawing == true &&
		g.BlockStates[nextBlock.Y][nextBlock.X] == Inactive {

		g.BlocksForScore++
		g.BlockStates[originalBlock.Y][originalBlock.X] = Exploding
		g.BlockStates[nextBlock.Y][nextBlock.X] = Exploding

		switch direction {
		case "up":
			// Not needed
		case "down":
			g.CheckScore("down", originalBlock, Pos{nextBlock.X, nextBlock.Y + 1})
		case "left":
			g.CheckScore("left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y})
		case "right":
			g.CheckScore("right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y})
		case "up_left":
			g.CheckScore("up_left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y - 1})
		case "up_right":
			g.CheckScore("up_right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y - 1})
		case "down_left":
			g.CheckScore("down_left", originalBlock, Pos{nextBlock.X - 1, nextBlock.Y + 1})
		case "down_right":
			g.CheckScore("down_right", originalBlock, Pos{nextBlock.X + 1, nextBlock.Y + 1})
		default:
		}
	}
}

// HandleScoreBlocks contains the logic for what should happen to blocks after they are marked by the CheckScore functions
func (g *GameBoard) HandleScoreBlocks() {
	if g.BlocksForScore >= 2 {
		g.BlockScorePausing = true
		g.LevelFallingTimer = 0
		g.SoundPlayer.PlaySound("break" + strconv.Itoa(1+rand.Intn(5)))
		g.DeGrayValue--
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
					if g.ScoreValue < g.MaxScoreValue {
						if g.ScoreValue+g.BlockPointValue < g.MaxScoreValue {
							g.ScoreValue += g.BlockPointValue
							g.LevelScoreValue += g.BlockPointValue
						} else {
							g.ScoreValue = g.MaxScoreValue
						}
					}
				}
			}
		}
	}

	for n := range g.BlockStates {
		for m := range g.BlockStates[n] {
			if g.Blocks[n][g.BlockStatesToGameBoard(m)].MainSprite.Drawing == true && g.BlockStates[n][m] != Inactive && g.BlockStates[n][m] != Active {
				g.BlockStates[n][m] = Inactive
			}
		}
	}

	g.BlocksForScore = 0
}
