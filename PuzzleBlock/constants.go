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
