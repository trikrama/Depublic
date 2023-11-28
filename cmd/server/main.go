package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/builder"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/binder"
	"github.com/trikrama/Depublic/internal/http/server"
	"github.com/trikrama/Depublic/internal/http/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	splash()

	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)

	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)

	runServer(srv, cfg.Port)

	waitForShutdown(srv)
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func splash() {
	colorReset := "\033[0m"

	splashText := `
	____        _                    
   |  _ \ _   _| |__   ___ _ __  ___ 
   | |_) | | | | '_ \ / _ \ '_ \/ __|
   |  __/| |_| | |_) |  __/ | | \__ \
   |_|    \__,_|_.__/ \___|_| |_|___/
									 
  `
	fmt.Println(colorReset, strings.TrimSpace(splashText))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
