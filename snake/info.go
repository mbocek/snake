package snake

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Info struct {
	//score int
	sizeX int
	sizeY int
}

const (
	infoWidth = 150
)

var (
	mplusNormalFont font.Face
	score           = 0
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewInfo(sizeX, sizeY int) *Info {
	i := &Info{
		sizeX: sizeX,
		sizeY: sizeY,
	}
	bus.Subscribe(topicFood, i.scoreHandler)
	return i
}

func (i *Info) Size() (int, int) {
	return i.sizeX, i.sizeY
}

func (i *Info) Update() error {
	return nil
}

func (i *Info) Draw(infoImage *ebiten.Image) {
	infoImage.Fill(color.RGBA{0xbb, 0xbb, 0x00, 0xaa})
	text.Draw(infoImage, strconv.Itoa(score), mplusNormalFont, 20, 40, color.White)
}

func (i *Info) scoreHandler(points int) {
	score += points
}
