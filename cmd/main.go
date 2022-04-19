package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sergey-timtsunyk/todo/pkg/handler"
	"github.com/sergey-timtsunyk/todo/pkg/infra"
	"github.com/sergey-timtsunyk/todo/pkg/repository"
	"github.com/sergey-timtsunyk/todo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initSettings()

	db, err := initDb()
	if err != nil {
		logrus.Fatalf("Error connection to mysql: %s", err.Error())
	}

	mongoDb, err := initMongo()
	if err != nil {
		logrus.Fatalf("Error connection to mongo: %s", err.Error())
	}

	repositories := repository.NewRepository(db, mongoDb)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(infra.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occurred while running server: %s", err.Error())
		}
	}()

	logrus.Print("Start todo app")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("End todo app")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error on server shut dowm: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error on db close: %s", err.Error())
	}

	if err := mongoDb.Client().Disconnect(context.Background()); err != nil {
		logrus.Fatalf("error on mongo disconect: %s", err.Error())
	}
}

func initSettings() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error env load: %s", err.Error())
	}
}

func initDb() (*sqlx.DB, error) {
	return repository.NewMysqlDB(repository.ConfigMysqlDB{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		DBName:   viper.GetString("mysql.db_name"),
		User:     viper.GetString("mysql.user"),
		Password: os.Getenv("DB_MYSQL_PASSWORD"),
	})
}

func initMongo() (*mongo.Database, error) {
	db, err := repository.NewConfigMongo(repository.ConfigMongo{
		Host:     viper.GetString("mongo.host"),
		Port:     viper.GetString("mongo.port"),
		DBName:   viper.GetString("mongo.db_name"),
		User:     viper.GetString("mongo.user"),
		Password: os.Getenv("DB_MONGO_PASSWORD"),
	})

	return db, err
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
