package assets

import (
	"image"
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestSpriteSheet_Get(t *testing.T) {
	// Create a mock sprite sheet (128x128 px, with 4x4 grid of 32x32 sprites)
	mockImage := ebiten.NewImage(128, 128)
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			mockImage.Set(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
		}
	}

	// Create the sprite sheet
	sheet := &spriteSheet{
		sheet:        mockImage,
		spriteWidth:  32,
		spriteHeight: 32,
		rows:         4,
		cols:         4,
	}

	// Test valid indices
	testCases := []struct {
		index int
		x0    int
		y0    int
		x1    int
		y1    int
	}{
		{index: 1, x0: 0, y0: 0, x1: 32, y1: 32},      // Top-left sprite
		{index: 5, x0: 0, y0: 32, x1: 32, y1: 64},     // First sprite in second row
		{index: 16, x0: 96, y0: 96, x1: 128, y1: 128}, // Bottom-right sprite
	}

	for _, tc := range testCases {
		t.Run("Valid index", func(t *testing.T) {
			sprite, err := sheet.Get(tc.index)
			assert.NoError(t, err)
			assert.NotNil(t, sprite)

			// Extract the bounds of the sprite to validate correctness
			subImage := sprite.SubImage(sprite.Bounds()).(*ebiten.Image)
			expected := mockImage.SubImage(image.Rect(tc.x0, tc.y0, tc.x1, tc.y1)).(*ebiten.Image)
			assert.Equal(t, expected.Bounds(), subImage.Bounds())
		})
	}

	// Test invalid indices
	invalidIndices := []int{-1, 0, 17} // Out of bounds for a 4x4 grid
	for _, idx := range invalidIndices {
		t.Run("Invalid index", func(t *testing.T) {
			sprite, err := sheet.Get(idx)
			assert.Error(t, err)
			assert.Nil(t, sprite)
		})
	}
}

type mockAssetManager struct{}

func (m *mockAssetManager) Get(path string) ([]byte, error) {
	return nil, nil
}

func (m *mockAssetManager) GetEbitenImage(path string) (*ebiten.Image, error) {
	// Return a mock image
	image := ebiten.NewImage(128, 128) // Mock image of 128x128 pixels
	return image, nil
}

func TestNewSpriteSheetGrid(t *testing.T) {
	mockManager := &mockAssetManager{}

	// Create a new sprite sheet
	spriteSheet, err := NewSpriteSheetGrid(mockManager, "path/to/mock.png", 32, 32, 4, 4)
	assert.NoError(t, err)
	assert.NotNil(t, spriteSheet)

	// Validate that the spriteSheet implements the SpriteSheet interface
	_, ok := spriteSheet.(SpriteSheet)
	assert.True(t, ok)
}
