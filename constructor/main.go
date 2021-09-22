package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/config"
	"github.com/vskurikhin/remote-sensing-platform/constructor/server"
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

	// Если логгирование на уровне трассировки включено, вывести
	// параметры конфигурации.
	if environ != nil && environ.Logging.Debug {
		fmt.Println(environ.String())
		logger.SetLevel(logger.DEBUG)
	}
	if environ != nil {
		logger.Printf("%s", environ.String())
	}
	// Создать инстанс сервера
	s := server.New(environ)
	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
