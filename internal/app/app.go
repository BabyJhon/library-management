package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	config "github.com/BabyJhon/library-managment/configs"
	"github.com/BabyJhon/library-managment/internal/handlers"
	"github.com/BabyJhon/library-managment/internal/repo"
	"github.com/BabyJhon/library-managment/internal/service"
	"github.com/BabyJhon/library-managment/pkg/httpserver"
	"github.com/BabyJhon/library-managment/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Run() {
	//logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Configs
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	//.env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	//DB
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("hui hui hui failed init db: %s", err.Error())
	}

	// Repositories
	repos := repo.NewRepository(db)

	// Service
	services := service.NewService(repos)

	// Handlers
	handlers := handlers.NewHandler(services)

	//HTTP server
	fmt.Println("start server")

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()

	logrus.Print("library-managment api started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error while server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error while data base closing: %s", err)
	}
}
