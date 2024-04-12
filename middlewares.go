package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/codelikesuraj/gdsc-challenge-day-nine-ten/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Middleware struct {
	DB *gorm.DB
}

func (m Middleware) IsAdmin(c *gin.Context) {
	tokenString, err := getBearerToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User

		err := m.DB.First(&user, claims["sub"]).Error
		switch {
		case err == gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid user"})
			return
		case err != nil:
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			return
		}

		// check if user is admin
		if !user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorised"})
			return
		}

		c.Next()

		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "bad"})
}

func (m Middleware) Authenticate(c *gin.Context) {
	tokenString, err := getBearerToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User

		err := m.DB.First(&user, claims["sub"]).Error
		switch {
		case err == gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid user"})
			return
		case err != nil:
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			return
		}

		c.Set("auth", user)

		c.Next()

		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "bad"})
}

func getBearerToken(c *gin.Context) (string, error) {
	log.Println(c.Request.Header)
	auth_header := c.Request.Header.Get("Authorization")
	log.Println(auth_header)
	if auth_header == "" {
		return "", errors.New("authorization header not set")
	}

	bearer_arr := strings.Split(strings.TrimSpace(auth_header), " ")
	if len(bearer_arr) != 2 {
		return "", errors.New("bearer token not set")
	}

	return bearer_arr[1], nil
}
