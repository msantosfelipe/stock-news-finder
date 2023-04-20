package config

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	UrlValorInveste  string `env:"URL_VALOR_INVESTE"`
	UrlBrazilJournal string `env:"URL_BRAZIL_JOURNAL"`
}

var ENV Environment

func init() {
	gotenv.Load() // load .env file (if exists)
	if _, err := env.UnmarshalFromEnviron(&ENV); err != nil {
		log.Fatal("Fatal error unmarshalling environment config: ", err)
	}
}
