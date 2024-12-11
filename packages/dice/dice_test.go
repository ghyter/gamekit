package dice

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

type mockSpriteSheet struct {
	sheet        *ebiten.Image
	spriteWidth  int
	spriteHeight int
	rows         int
	cols         int
}

func (m *mockSpriteSheet) Sheet() *ebiten.Image {
	return m.sheet
}

func (m *mockSpriteSheet) Draw(screen *ebiten.Image) {

}

func (m *mockSpriteSheet) Get(index int) (*ebiten.Image, error) {
	// Return a mock image for valid indices, error for invalid indices
	if index < 1 || index > m.rows*m.cols {
		return nil, assert.AnError
	}
	img := ebiten.NewImage(m.spriteWidth, m.spriteHeight) // Create a blank image
	return img, nil
}

func TestNewDie(t *testing.T) {
	// Test valid die creation
	die, err := NewDie(D6)
	assert.NoError(t, err)
	assert.NotNil(t, die)

	// Test that die returns correct type and face count
	d6, ok := die.(*dieImpl)
	assert.True(t, ok)
	assert.Equal(t, D6, d6.diceType)
	assert.Equal(t, 6, d6.sides)
}

func TestDieImpl_Roll(t *testing.T) {
	// Mock SpriteSheet
	mockSheet := &mockSpriteSheet{
		spriteWidth:  32,
		spriteHeight: 32,
		rows:         8,
		cols:         8,
	}

	// Create a D6 die
	die := &dieImpl{
		diceType: D6,
		sheet:    mockSheet,
		sides:    6,
	}

	// Test rolling the die
	rolledFace, img := die.Roll()
	assert.GreaterOrEqual(t, rolledFace, 0) // Check roll range
	assert.Less(t, rolledFace, 6)
	assert.NotNil(t, img) // Ensure an image is returned
}

func TestDieImpl_Set(t *testing.T) {
	// Mock SpriteSheet
	mockSheet := &mockSpriteSheet{
		spriteWidth:  32,
		spriteHeight: 32,
		rows:         8,
		cols:         8,
	}

	// Create a D6 die
	die := &dieImpl{
		diceType: D6,
		sheet:    mockSheet,
		sides:    6,
	}

	// Test setting the face to a valid value
	img, err := die.Set(3)
	assert.NoError(t, err)
	assert.NotNil(t, img)

	// Test setting the face to an invalid value
	img, err = die.Set(7)
	assert.Error(t, err)
	assert.Nil(t, img)
}

func TestDieEdgeCases(t *testing.T) {
	// Test cases for each DieType
	testCases := []struct {
		dieType     DieType
		faceCount   int
		spriteStart int
	}{
		{D4, 4, 57},
		{D6, 6, 1},
		{D6Alt, 6, 9},
		{D8, 8, 17},
		{D10, 10, 10},
		{D12, 12, 25},
		{D20, 20, 37},
	}

	for _, tc := range testCases {
		t.Run(tc.dieType.String(), func(t *testing.T) {
			// Create a new die
			die, err := NewDie(tc.dieType)
			assert.NoError(t, err)
			assert.NotNil(t, die)

			// Test Roll method
			for i := 0; i < 100; i++ { // Roll multiple times to test randomness
				rolledFace, img := die.Roll()
				assert.GreaterOrEqual(t, rolledFace, 0)
				assert.LessOrEqual(t, rolledFace, tc.faceCount)
				assert.NotNil(t, img)
			}

			// Test Set method
			for face := 1; face <= tc.faceCount; face++ {
				img, err := die.Set(face)
				assert.NoError(t, err)
				assert.NotNil(t, img)
			}

			// Test invalid Set values
			invalidFaces := []int{-1, 0, tc.faceCount + 1, tc.faceCount + 5}
			for _, face := range invalidFaces {
				img, err := die.Set(face)
				assert.Error(t, err)
				assert.Nil(t, img)
			}
		})
	}
}
