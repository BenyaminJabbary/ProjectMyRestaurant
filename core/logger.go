package core

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func NewFileLogger() *zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Logger.Println("failed to open log file: ", err)
	}

	logger := zerolog.New(file).
		With().
		Timestamp().
		Str("AppName", "Restaurant").
		Logger()


	return &logger

}
