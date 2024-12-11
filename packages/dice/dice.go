package dice

import (
	"fmt"
	"time"

	"github.com/ghyter/gamekit/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
)

type Die interface {
	Roll() (int, *ebiten.Image)
	Set(int) (*ebiten.Image, error)
	LastRoll() (int, *ebiten.Image)
	Type() DieType
	Sheet() assets.SpriteSheet
}

type DieType int

const (
	D4 DieType = iota
	D6
	D6Alt
	D8
	D10
	D12
	D20
)

func (d DieType) String() string {
	switch d {
	case D4:
		return "D4"
	case D6:
		return "D6"
	case D6Alt:
		return "D6"
	case D8:
		return "D8"
	case D10:
		return "D10"
	case D12:
		return "D12"
	case D20:
		return "D20"
	default:
		return "D6"
	}
}

func (d DieType) FaceCount() int {
	switch d {
	case D4:
		return 4
	case D6:
		return 6
	case D6Alt:
		return 6
	case D8:
		return 8
	case D10:
		return 10
	case D12:
		return 12
	case D20:
		return 20
	default:
		return 6
	}
}

func (d DieType) SpriteStart() int {
	switch d {
	case D4:
		return 57
	case D6:
		return 1
	case D6Alt:
		return 9
	case D8:
		return 17
	case D10:
		return 25
	case D12:
		return 25
	case D20:
		return 37
	default:
		return 1
	}
}

func NewDie(d DieType) (Die, error) {
	am := assets.NewAssetManager(assets.AssetTypeDice)
	sp, err := assets.NewSpriteSheetGrid(am, "dice/dice.png", 32, 32, 8, 8)

	if err != nil {
		return nil, err
	}
	rand.Seed(uint64(time.Now().UnixNano()))

	return &dieImpl{
		diceType: d,
		sides:    d.FaceCount(),
		sheet:    sp,
		lastRoll: 0,
	}, nil
}

type dieImpl struct {
	diceType DieType
	sheet    assets.SpriteSheet
	sides    int
	lastRoll int
}

func (d *dieImpl) Type() DieType {
	return d.diceType
}

func (d *dieImpl) Sheet() assets.SpriteSheet {
	return d.sheet
}

// Roll will select a random number and return it, and the image associated with it.
func (d *dieImpl) Roll() (int, *ebiten.Image) {
	face := rand.Intn(d.sides)
	d.lastRoll = face + 1
	spriteStart := d.diceType.SpriteStart()
	spriteIndex := face + spriteStart - 1
	fmt.Printf("Rolling %s: value: %d,getSprite: %d\n", d.diceType.String(), d.lastRoll, spriteStart)
	img, err := d.sheet.Get(spriteIndex)
	if err != nil {
		return face, nil
	}

	return d.lastRoll, img
}

// Set will return the image for a given value.  This is the index of the sprite.

func (d *dieImpl) Set(face int) (*ebiten.Image, error) {

	if face <= 0 || face > d.sides {
		return nil, fmt.Errorf("face value %d is out of range for %s", face, d.diceType)
	}

	d.lastRoll = face + 1
	img, err := d.sheet.Get(face + d.diceType.SpriteStart())
	return img, err

}

// Roll will select a random number and return it, and the image associated with it.
func (d *dieImpl) LastRoll() (int, *ebiten.Image) {
	img, err := d.sheet.Get(d.lastRoll + d.diceType.SpriteStart())
	if err != nil {
		return d.lastRoll, nil
	}

	return d.lastRoll + 1, img
}
