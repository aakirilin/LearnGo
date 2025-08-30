package mock

import (
	"Server/dto"
)

var TestUsers = []dto.UserDTO{
	{1, "admin@test.ru", "admin"},
	{2, "user@test.ru", "user"},
}

var TestTasks = []dto.TaskDTO{
	{1, "admin test", 1, 1},
	{2, "user test", 2, 2},
}
