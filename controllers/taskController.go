package controllers

import (
	"Server/dto"
	"Server/mock"
	"net/http"

	"github.com/gin-gonic/gin"
)

func maxTaskIndex(tasks []dto.TaskDTO) int {
	res := 0
	for _, t := range tasks {
		if t.Id > res {
			res = t.Id
		}
	}
	return res
}

type TaskController struct {
	users []dto.UserDTO
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, mock.TestTasks)
}

func (tc *TaskController) AddTasks(c *gin.Context) {
	var newTask dto.TaskDTO
	if c.BindJSON(&newTask) == nil {
		newTask.Id = maxTaskIndex(mock.TestTasks) + 1
		mock.TestTasks = append(mock.TestTasks, newTask)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача добавлена",
	})
}
