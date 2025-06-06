package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	consoleWriter.TimeFormat = "15:04:05"

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger().Level(zerolog.InfoLevel)
}
