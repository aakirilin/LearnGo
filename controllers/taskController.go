package controllers

import (
	dto "Server/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

var testTasks = []dto.TaskDTO{
	{1, "admin test", 1, 1},
	{2, "user test", 2, 2},
}

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
	c.JSON(http.StatusOK, testTasks)
}

func (tc *TaskController) AddTasks(c *gin.Context) {
	var newTask dto.TaskDTO
	if c.BindJSON(&newTask) == nil {
		newTask.Id = maxTaskIndex(testTasks) + 1
		testTasks = append(testTasks, newTask)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача добавлена",
	})
}
