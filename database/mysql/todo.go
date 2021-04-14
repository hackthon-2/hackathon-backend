package database

import (
	"hackthon/database"
	"hackthon/model"
)

func CreateToDo(todo *model.Todo) (error, int64) {
	result := database.DB().Table("todo").Create(todo)
	return result.Error, result.RowsAffected
}
