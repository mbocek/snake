package main

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mbocek/snake/snake"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{DisableColors: true, FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	game := snake.NewGame(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
