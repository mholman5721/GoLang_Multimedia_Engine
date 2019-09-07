package main

import (
	"golang-games/PuzzleBlock/gameboard"

	"github.com/veandco/go-sdl2/sdl"
)

type mouseState struct {
	leftButton  bool
	rightButton bool
	x, y        int
}

var currentMouseState mouseState
var prevMouseState mouseState

func getMouseState() mouseState {
	mouseX, mouseY, mouseButtonState := sdl.GetMouseState()
	leftButton := mouseButtonState & sdl.ButtonLMask()
	rightButton := mouseButtonState & sdl.ButtonRMask()

	var result mouseState

	result.x = int(mouseX)
	result.y = int(mouseY)
	result.leftButton = !(leftButton == 0)
	result.rightButton = !(rightButton == 0)

	return result
}

var keyboardState []uint8
var prevKeyboardState []uint8

// KeyDownOnce returns true if the key has been pressed once
func KeyDownOnce(key uint8) bool {
	return keyboardState[key] == 1 && prevKeyboardState[key] == 0
}

// KeyPressed returns true if the key is currently pressed
func KeyPressed(key uint8) bool {
	return keyboardState[key] == 0 && prevKeyboardState[key] == 1
}

func getKeyboardState(g *gameboard.GameBoard) {
	if sdl.GetKeyboardFocus() == window || sdl.GetMouseFocus() == window {

		if KeyDownOnce(sdl.SCANCODE_UP) {
			g.MoveActiveBlock("up")
		}
		if KeyDownOnce(sdl.SCANCODE_DOWN) {
			g.MoveActiveBlock("down")
		}
		if KeyDownOnce(sdl.SCANCODE_LEFT) {
			g.MoveActiveBlock("left")
		}
		if KeyDownOnce(sdl.SCANCODE_RIGHT) {
			g.MoveActiveBlock("right")
		}

		for i, v := range keyboardState {
			prevKeyboardState[i] = v
		}
	}
}

func initInput() {
	keyboardState = sdl.GetKeyboardState()
	prevKeyboardState = make([]uint8, len(keyboardState))
	for i, v := range keyboardState {
		prevKeyboardState[i] = v
	}
	currentMouseState = getMouseState()
	prevMouseState = currentMouseState
}
