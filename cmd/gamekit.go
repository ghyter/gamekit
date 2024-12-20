package main

import (
	"fmt"
	"image/color"

	"github.com/ghyter/gamekit/packages/dice"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game interface {
	ebiten.Game
}

type DiceGame struct {
	dieTypes []dice.DieType
	dice     []dice.Die
}

func (g *DiceGame) Update() error {

	for i, d := range g.dice {
		if inpututil.IsKeyJustPressed(ebiten.Key1 + ebiten.Key(i)) {
			val, _ := d.Roll()
			fmt.Printf("Die %s: Rolled %d\n", d.Type().String(), val)
		}
	}

	return nil
}

func (g *DiceGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xcc, 0xcc, 0xcc, 0x00})
	g.dice[0].Sheet().Draw(screen)

	for i, d := range g.dice {
		_, lastRoll := d.LastRoll()
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64((i+1)*35), 700)
		screen.DrawImage(lastRoll, opts)
	}

	

}

func (g *DiceGame) Layout(outsideWidth int, outsideHeight int) (screenX int, screenY int) {

	return outsideWidth, outsideHeight
}

func main() {

	dg := &DiceGame{
		dieTypes: []dice.DieType{dice.D4, dice.D6, dice.D6Alt, dice.D8, dice.D12, dice.D20},
	}
	for _, dieType := range dg.dieTypes {
		d, err := dice.NewDie(dieType)
		if err != nil {
			panic(err)
		}
		dg.dice = append(dg.dice, d)
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Dice Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.RunGame(dg)
}
