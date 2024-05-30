package main

import (
	"AuthService/data"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func AuthServer() *Server {
	router := gin.Default()
	var dsn string = "host=localhost user=postgres password=titbandau dbname=auth port=5432 sslmode=disable"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&data.User{})
	if err != nil {
		log.Fatal("Connecting to database failed")
	}

	return &Server{
		router: router,
		db:     db,
	}
}

func (s *Server) Run() error {

	c := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	s.router.Use(c)
	s.addAuth()
	return s.router.Run(":9000")
}
