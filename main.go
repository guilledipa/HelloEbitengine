package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type game struct {
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "¡Hola mundo!")
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrint(screen, "Clicked!")
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrint(screen, "Clicked!")
	}
	if IsClicked() {
		ebitenutil.DebugPrint(screen, "Funcion IsClicked!")
	}
	// Mouse coords
	s := fmt.Sprintln(ebiten.CursorPosition())
	ebitenutil.DebugPrint(screen, s)
	// Keyboard
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		ebitenutil.DebugPrint(screen, "A is pressed")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		ebitenutil.DebugPrint(screen, "B is just pressed")
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func main() {
	ebiten.SetWindowTitle("Ebitengine ~ ¡Hola mundo!")
	ebiten.SetWindowSize(640, 480)
	g := &game{}
	ebiten.RunGame(g)
}

// IsClicked Detects mouse button presses as well as touches.
// Since touch can occur in multiple places at the same time,
// ebiten.AppendTouchIDsreturns a slice of IDs that identify fingers touching
// the touch panel. If the number of elements in the returned value is not zero,
// it means that there is a finger touching the touch panel, so returns true.
func IsClicked() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return len(ebiten.AppendTouchIDs(nil)) != 0
}
