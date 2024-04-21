package main

import (
	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Broker(c *gin.Context) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit the Broker",
	}
	c.Header("Content-Type", "application/json")
	c.JSON(200, payload)

}

func (s *Server) addBroker(router *gin.Engine) {
	router.POST("/broker", Broker)
}
