package main

import (
	"AuthService/data"
	"AuthService/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

func (s *Server) CreateUser(c *gin.Context) {
	var req RegisterRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	user := data.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  string(hashedPassword),
	}
	id, err := repository.Insert(user, s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (s *Server) GetUsers(c *gin.Context) {
	users, err := repository.GetAll(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(loginRequest)
	user, err := repository.GetByEmail(loginRequest.Email, s.db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	match, err := repository.PasswordMatches(user.Password, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) addAuth() {
	s.router.POST("/auth/login", s.Login)
	s.router.POST("/auth/register", s.CreateUser)
	s.router.GET("/auth/users", s.GetUsers)

}
