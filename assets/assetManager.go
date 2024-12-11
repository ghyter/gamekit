package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed dice
var diceAssets embed.FS

type AssetManager interface {
	Get(path string) ([]byte, error)
	GetEbitenImage(path string) (*ebiten.Image, error)
}

type assetManager struct {
	cache  map[string][]byte
	assets embed.FS
}

func (a *assetManager) Get(path string) ([]byte, error) {
	if data, ok := a.cache[path]; ok {
		return data, nil
	}
	data, err := a.assets.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load asset %s: %w", path, err)
	}
	a.cache[path] = data
	return data, nil
}

func (a *assetManager) GetEbitenImage(path string) (*ebiten.Image, error) {
	rawimage, err := a.Get(path)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(rawimage)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil

}

type AssetType int

const (
	AssetTypeDice AssetType = iota
)

func NewAssetManager(assettype AssetType) AssetManager {

	var lfs embed.FS
	switch assettype {
	case AssetTypeDice:
		lfs = diceAssets
	}

	return &assetManager{
		cache:  make(map[string][]byte),
		assets: lfs,
	}

}
