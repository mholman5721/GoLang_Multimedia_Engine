package gameboard

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
			//g.ProccessBlockMovement("Y--")
		case "down":
			//fmt.Println("down")
			g.ProccessBlockMovement("Y++")
			g.LevelFallingTimer = 0
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
