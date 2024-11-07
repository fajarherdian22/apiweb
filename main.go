package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fajarherdian22/apiweb/controller"
	"github.com/fajarherdian22/apiweb/db"
	"github.com/fajarherdian22/apiweb/middleware"
	"github.com/fajarherdian22/apiweb/repository"
	"github.com/fajarherdian22/apiweb/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("configs")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	config_db := viper.GetStringMapString("databases.db_credential")
	dbcon := db.ConDB(config_db)

	queries := repository.New(dbcon)
	dataService := service.NewService(queries, dbcon)
	dataController := controller.NewDataController(dataService)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
		AllowHeaders:    []string{"Content-Type", "Origin"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	// db.NewRedisClient(viper.GetString("redis.host"), viper.GetString("redis.password"))
	// router.Use(middleware.RedisMiddleware())
	router.Use(middleware.LoggingMiddleware())
	router.Use(gin.Recovery())

	r := router.Group("/api/")
	{
		r.POST("/data/city", dataController.GetCityLink)
		r.GET("/data/list/city", dataController.ListCity)
		r.POST("topo/city", dataController.GetCityTopo)
		r.GET("/topo/list/city", dataController.ListCityTopo)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
