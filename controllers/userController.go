package controllers

import (
	"Server/mock"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (lc *UserController) GetUser(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ид пользователя должно быть числом",
		})
		return
	}
	for _, u := range mock.TestUsers {
		if u.Id == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Пользователь с таким Ид не найден",
	})
}
