package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/rtexty/AutoFreelance/config"
)

const(
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main(){

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
	
	// TODO: 3. Инициализация хранилища (database)
	
	// TODO: 4. Инициализация роутера (chi/render)

	// TODO: 5. Запуск сервера

	fmt.Println("Hello, world!")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}