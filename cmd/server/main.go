package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"github.com/NhatPixel/cinema-service/config"
	"github.com/NhatPixel/cinema-service/internal/handler"
	"github.com/NhatPixel/cinema-service/internal/repository"
	"github.com/NhatPixel/cinema-service/internal/service"
	appvalidator "github.com/NhatPixel/cinema-service/internal/validator"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		appvalidator.RegisterCinemaValidation(v)
	}

	db, err := config.NewMySQL()
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	cinemaRepo := repository.NewCinemaRepo(db)
	cinemaService := service.NewCinemaService(cinemaRepo)
	cinemaHandler := handler.NewCinemaHandler(cinemaService)

	r.POST("/cinemas", cinemaHandler.Create)
	r.GET("/cinemas", cinemaHandler.Get)
	r.PUT("/cinemas", cinemaHandler.Update)
	r.DELETE("/cinemas/:id", cinemaHandler.Delete)

	r.Run(":8080")
}
