package assets

import (
	_ "image/png" // Register PNG decoder
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetManager_Get(t *testing.T) {
	// Initialize AssetManager with real dice assets
	manager := &assetManager{
		cache:  make(map[string][]byte),
		assets: diceAssets,
	}

	// Test loading an existing asset
	data, err := manager.Get("dice/dice.png")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)

	// Test caching by requesting the same asset again
	cachedData, err := manager.Get("dice/dice.png")
	assert.NoError(t, err)
	assert.Equal(t, data, cachedData)

	// Test loading a non-existent asset
	_, err = manager.Get("dice/nonexistent.png")
	assert.Error(t, err)
}

func TestAssetManager_GetEbitenImage(t *testing.T) {
	// Initialize AssetManager with real dice assets
	manager := &assetManager{
		cache:  make(map[string][]byte),
		assets: diceAssets,
	}

	// Test loading and decoding a real image
	img, err := manager.GetEbitenImage("dice/dice.png")
	assert.NoError(t, err)
	assert.NotNil(t, img)
	assert.Equal(t, img.Bounds().Dx(), 256) // Replace with the actual width of dice.png
	assert.Equal(t, img.Bounds().Dy(), 256) // Replace with the actual height of dice.png

	// Test loading a non-image file (e.g., SVG or invalid asset)
	_, err = manager.GetEbitenImage("dice/dice.svg")
	assert.Error(t, err)

	// Test loading a non-existent asset
	_, err = manager.GetEbitenImage("dice/nonexistent.png")
	assert.Error(t, err)
}

func TestNewAssetManager(t *testing.T) {
	// Test creating an AssetManager for dice assets
	manager := NewAssetManager(AssetTypeDice)
	assert.NotNil(t, manager)

}
