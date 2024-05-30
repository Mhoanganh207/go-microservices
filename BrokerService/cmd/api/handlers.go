package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActionRequest struct {
	Action      string      `json:"action"`
	AuthRequest AuthRequest `json:"auth"`
}

func (s *Server) Broker(c *gin.Context) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit the Broker",
	}
	c.Header("Content-Type", "application/json")
	c.JSON(200, payload)

}

func (s *Server) Submission(c *gin.Context) {

	var req ActionRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch req.Action {
	case "auth":
		executeAuth(c, req.AuthRequest)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
	}

}

func executeAuth(c *gin.Context, req AuthRequest) {
	requestBody, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request, err := http.NewRequest("POST", "http://localhost:9000/auth/login", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	switch response.StatusCode {
	case 200:
		c.JSON(http.StatusOK, gin.H{"message": "User authenticated"})
	case 401:
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	case 500:
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	}

}

func (s *Server) addBroker() {
	s.router.POST("/broker", s.Broker)
	s.router.POST("/broker/submission", s.Submission)
}
