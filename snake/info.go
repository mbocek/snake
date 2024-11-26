package snake

import (
	"github.com/rs/zerolog/log"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Info struct {
	score           int
	sizeX           int
	sizeY           int
	speed           *Speed
	mplusNormalFont font.Face
}

const (
	infoWidth = 150
)

func (i *Info) init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load font")
	}

	const dpi = 72
	i.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot open font")
	}
}

func NewInfo(sizeX, sizeY int, speed *Speed) *Info {
	i := &Info{
		sizeX: sizeX,
		sizeY: sizeY,
		speed: speed,
	}
	if err := bus.Subscribe(topicFood, i.scoreHandler); err != nil {
		log.Fatal().Err(err).Str("topic", topicFood).Msg("cannot subscribe to topic")
	}
	i.init()
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
	text.Draw(infoImage, "Score: "+strconv.Itoa(i.score), i.mplusNormalFont, 20, 40, color.White)
	text.Draw(infoImage, "Speed: "+strconv.Itoa(i.speed.Speed()), i.mplusNormalFont, 20, 80, color.White)
}

func (i *Info) scoreHandler(points int) {
	i.score += points
}
