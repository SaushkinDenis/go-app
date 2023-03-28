package main

import (
	todo "github.com/SaushkinDenis/go-app"
	"github.com/SaushkinDenis/go-app/pkg/handler"
	"github.com/SaushkinDenis/go-app/pkg/repository"
	"github.com/SaushkinDenis/go-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initialization config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to ibnitialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if serverError := server.Run(viper.GetString("port"), handlers.InitRoutes()); serverError != nil {
		logrus.Fatalf("error occured while running server: %s", serverError.Error())
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	return viper.ReadInConfig()
}
