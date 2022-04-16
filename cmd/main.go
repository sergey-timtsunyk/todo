package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sergey-timtsunyk/todo/pkg/handler"
	"github.com/sergey-timtsunyk/todo/pkg/infra"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
	"github.com/sergey-timtsunyk/todo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error env load: %s", err.Error())
	}

	db, err := repository.NewMysqlDB(repository.ConfigMysqlDB{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		DBName:   viper.GetString("mysql.db_name"),
		User:     viper.GetString("mysql.user"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Error connection to mysql: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(infra.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occurred while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
