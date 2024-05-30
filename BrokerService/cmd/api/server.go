package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	router := gin.Default()
	return &Server{
		router: router,
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
	s.addBroker()
	return s.router.Run(":8000")
}
