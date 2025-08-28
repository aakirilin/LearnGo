package dto

type TaskDTO struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	AutiorId int    `json:"autorId"`
	ExuterId int    `json:"exuterId"`
}
