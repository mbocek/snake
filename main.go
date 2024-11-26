package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mbocek/snake/snake"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	game := snake.NewGame(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal().Err(err).Msg("cannot run the game")
	}
}
