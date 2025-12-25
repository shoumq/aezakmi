package main

import (
	"aezakmi/internal/config"
	"aezakmi/internal/handler"
	"aezakmi/internal/middleware"
	"aezakmi/internal/model"
	"aezakmi/internal/repository"
	"aezakmi/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("starting Agora service")

	cfg := config.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		sugar.Fatalw("failed to connect to database", "error", err)
	}

	err = db.AutoMigrate(&model.Room{}, &model.TokenLog{})
	if err != nil {
		sugar.Fatalw("failed to migrate database", "error", err)
	}
	repo := repository.NewRepository(db, cfg, sugar)
	svc := service.NewService(repo, sugar)
	h := handler.NewHandler(svc, cfg, sugar)

	r := gin.Default()
	api := r.Group("/api")
	api.Use(middleware.SimpleAuth())

	api.POST("/auth/simple", h.AuthSimple)
	api.POST("/tokens/rtc", h.GenerateRTCToken)
	api.POST("/tokens/rtm", h.GenerateRTMToken)
	api.POST("/rooms", h.CreateRoom)

	sugar.Infow("server running on :8080")
	if err := r.Run(":8080"); err != nil {
		sugar.Fatalw("failed to run server", "error", err)
	}
}
