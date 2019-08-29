package gameboard

import "math/rand"

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
					if g.ScoreValue < g.MaxScoreValue {
						if g.ScoreValue+g.BlockPointValue < g.MaxScoreValue {
							g.ScoreValue += g.BlockPointValue
							g.LevelScoreValue += g.BlockPointValue
							g.DeGrayValue--
						} else {
							g.ScoreValue = g.MaxScoreValue
							g.DeGrayValue--
						}
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
