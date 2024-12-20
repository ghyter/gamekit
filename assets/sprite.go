package assets

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

/*


 */
// Sprite indexes start at 1, not 0.
// In this chart, the zeros are for spacing.
// 01 02 03 04 05 06 07 08
// 09 10 11 12 13 14 15 16
// 17 18 19 20 21 22 23 24
// 25 26 27 28 29 30 31 32
// 33 34 35 36 37 38 39 40
// 41 42 43 44 45 46 47 48
// 49 50 51 52 53 54 55 56
// 57 58 59 60 61 62 63 64

type SpriteSheet interface {
	Get(int) (*ebiten.Image, error)
	Draw(screen *ebiten.Image)
}

type spriteSheet struct {
	sheet        *ebiten.Image
	spriteWidth  int
	spriteHeight int
	rows         int
	cols         int
}

/*
Get will return the specific sprite found at the given location.
The index is 1-based, not 0-based.
*/
func (s *spriteSheet) Get(index int) (*ebiten.Image, error) {
	if index < 1 || index > s.rows*s.cols {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}

	//fmt.Printf("Get Sprite: %d\n", index)

	// Convert 1-based index to 0-based row and column
	index -= 1

	row := index / s.cols
	col := index % s.cols
	//fmt.Printf("Get Sprite: Position (%d,%d)\n", row, col)

	// Calculate the position of the sprite in the sheet
	x0 := col * s.spriteWidth
	y0 := row * s.spriteHeight
	x1 := x0 + s.spriteWidth
	y1 := y0 + s.spriteHeight
	//fmt.Printf("Get Sprite: Upper Left (%d, %d), Lower Right (%d,%d)\n", x0, y0, x1, y1)

	// Extract and return the sprite as a sub-image
	return s.sheet.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), nil
}

/*
Draw will render all sprites on the screen with their indices below.
*/
func (s *spriteSheet) Draw(screen *ebiten.Image) {
	for i := 1; i <= s.rows*s.cols; i++ {
		// Get the sprite
		sprite, err := s.Get(i)
		if err != nil {
			fmt.Printf("Error retrieving sprite %d: %v\n", i, err)
			continue
		}

		// Calculate the screen position
		row := (i - 1) / s.cols
		col := (i - 1) % s.cols
		x := float64(col * (s.spriteWidth + 10))  // Add padding between sprites
		y := float64(row * (s.spriteHeight + 20)) // Add padding for text

		// Draw the sprite
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(x, y)
		screen.DrawImage(sprite, opts)

		// Draw the index below the sprite
		textX := x + float64(s.spriteWidth)/2 - 10 // Center text under sprite
		textY := y + float64(s.spriteHeight) + 10
		text.Draw(screen, fmt.Sprintf("%02d", i), basicfont.Face7x13, int(textX), int(textY), color.White)
	}
}

func NewSpriteSheetGrid(assetManager AssetManager, path string, spriteWidth, spriteHeight, rows, cols int) (SpriteSheet, error) {

	image, err := assetManager.GetEbitenImage(path)
	if err != nil {
		return nil, err
	}
	return &spriteSheet{
		sheet:        image,
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		rows:         rows,
		cols:         cols,
	}, nil
}
