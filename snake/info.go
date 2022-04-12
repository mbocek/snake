package snake

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Info struct {
	//score int
	sizeX int
	sizeY int
}

const (
	infoWidth = 150
)

func NewInfo(sizeX, sizeY int) *Info {
	i := &Info{
		sizeX: sizeX,
		sizeY: sizeY,
	}
	return i
}

func (i *Info) Size() (int, int) {
	return i.sizeX, i.sizeY
}

func (i *Info) Update(inout *Input) error {
	return nil
}

func (i *Info) Draw(infoImage *ebiten.Image) {
	infoImage.Fill(color.RGBA{0xbb, 0xbb, 0x00, 0xaa})
}
