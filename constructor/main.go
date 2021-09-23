package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/config"
	"github.com/vskurikhin/remote-sensing-platform/constructor/server"
	"github.com/vskurikhin/remote-sensing-platform/constructor/server/handlers"
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
	// Создать инстанс сервера
	s := server.New(environ)
	// Обработчики запросов.
	h := handlers.Handlers{Server: s}
	// Зарегистрировать маршруты для поиска пользователей.
	s.GET("/api/admin/{userId}/polls/{pollId}/constructor", h.GetPollConstructorPageData)
	// Run
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
