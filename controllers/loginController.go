package controllers

import (
	"Server/dto"
	"Server/mock"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

type Claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginController struct {
}

func (lc *LoginController) generateJWT(user dto.UserDTO) (string, error) {
	expirationTime := time.Now().Add(160 * time.Minute)
	claims := &Claims{
		Id:    user.Id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)

}

func (lc *LoginController) Login(c *gin.Context) {

	var user dto.UserDTO
	if c.BindJSON(&user) == nil {
		for _, u := range mock.TestUsers {
			if u.Email == user.Email && u.Password == user.Password {
				token, _ := lc.generateJWT(u)
				tokens = append(tokens, token)

				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Такого пользователя не нашли",
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Что то не так с данными",
	})
}

// Middleware для ограничения запросов
func (lc *LoginController) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		tokenString := strings.Split(bearerToken, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return jwtKey, nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Что то не так с токеном",
			})
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok {
			c.Next()
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Пользователь не авторизован",
			})
			return
		}
	}
}
