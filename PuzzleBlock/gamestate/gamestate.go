package gamestate

// GameState denotes which phase of the game is being updated
type GameState int

const (
	// TitleScreen is the first screen the player sees
	TitleScreen GameState = iota
	// OptionsScreen allows the player to set various options like sound volume, number of levels, etc.
	OptionsScreen
	// MainGame is where the game is actually played
	MainGame
	// Transitioning is an intermediate state that transfers from one state to another
	Transitioning
	// QuitGame exits the game
	QuitGame
)
