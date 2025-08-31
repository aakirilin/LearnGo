package controllers

import (
	"Server/dto"
	"Server/mock"
	sseSever "Server/sse"
	"net/http"
	"strconv"

	"github.com/alexandrevicenzi/go-sse"
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

func (tc *TaskController) GetTask(c *gin.Context) {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ид задачи должно быть числом",
		})
		return
	}
	for _, t := range mock.TestTasks {
		if t.Id == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Задача с таким Ид не найдена",
	})
}

func (tc *TaskController) AddTasks(c *gin.Context) {
	var newTask dto.TaskDTO
	if c.BindJSON(&newTask) == nil {
		newTask.Id = maxTaskIndex(mock.TestTasks) + 1
		mock.TestTasks = append(mock.TestTasks, newTask)
	}

	go sseSever.SseServer.SendMessage("/events/addtasks", sse.SimpleMessage(strconv.Itoa(newTask.Id)))

	c.JSON(http.StatusOK, gin.H{
		"message": "Задача добавлена",
	})
}
