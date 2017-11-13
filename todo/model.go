package todo

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	Message string `json:"message,omitempty"`
}

func (Model) TableName() string {
	return "todos"
}