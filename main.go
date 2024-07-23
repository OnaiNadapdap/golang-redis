package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/onainadapdap1/golang-crud-redis/config"
	"github.com/onainadapdap1/golang-crud-redis/controller"
	"github.com/onainadapdap1/golang-crud-redis/database"
	"github.com/onainadapdap1/golang-crud-redis/model"
	"github.com/onainadapdap1/golang-crud-redis/repo"
	"github.com/onainadapdap1/golang-crud-redis/router"
	"github.com/onainadapdap1/golang-crud-redis/usecase"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello world sarah!")
	// mysql setup

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load environment variables", err)
	}

	db := database.ConnectionMySQLDB(&loadConfig)
	db.AutoMigrate(&model.Novel{})

	rdb := database.ConnectionRedisDB(&loadConfig)
	startServer(db, rdb)
	
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()
	novelRepo := repo.NewNovelRepo(db, rdb)
	novelUseCase := usecase.NewNovelUseCase(novelRepo)
	novelController := controller.NewNovelController(novelUseCase)

	routes := router.NewRouter(app, novelController)

	err := routes.Listen(":3400")
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected at port :3400")
	
}