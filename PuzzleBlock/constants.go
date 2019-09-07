package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// WinWidth denotes the width of the window
const WinWidth int = 1280

// WinHeight denotes the height of the window
const WinHeight int = 720

// WinDepth denotes the 'depth' of the window for pseudo 3d effects
const WinDepth int = 100

var window *sdl.Window
var renderer *sdl.Renderer

// GameState denotes which phase of the game is being updated
type GameState int

const (
	// TitleScreen is the first screen the player sees
	TitleScreen GameState = iota
	// OptionsScreen allows the player to set various options like sound volume, number of levels, etc.
	OptionsScreen
	// MainGame is where the game is actually played
	MainGame
)
