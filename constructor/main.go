package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/config"
)

func main() {
	// Загрузка конфигурации
	var envFile string
	flag.StringVar(&envFile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()
	err := godotenv.Load(envFile)

	if err != nil {
		logger.Debug("main: can't load configuration")
	}
	environ, err := config.Environ()
	if environ != nil {
		logger.Printf("%s", environ.String())
	}
}
