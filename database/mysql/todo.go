package database

import (
	"hackthon/database"
	"hackthon/model"
	"time"
)

// CreateToDo 创建待办数据
func CreateToDo(todo *model.Todo) (error, int64) {
	result := database.DB().Table("todo").Create(todo)
	return result.Error, result.RowsAffected
}

// UpdateToDo 更新待办数据,根据拉取列表时获取到的id来更新
func UpdateToDo(todoID uint, todo *model.Todo) (error, int64) {
	result := database.DB().Table("todo").Where("id=?", todoID).Updates(todo)
	return result.Error, result.RowsAffected
}

func FindToDoById(todoID uint) (model.Todo, int64, error) {
	var todo model.Todo
	result := database.DB().Table("todo").Where("id=?", todoID).First(&todo)
	if result.RowsAffected != 1 {
		return todo, result.RowsAffected, result.Error
	}
	todo.Time = todo.Time[:10]
	return todo, result.RowsAffected, result.Error
}

func ListTodoByTime(userId uint, date string) ([]model.Todo, error) {
	var todo []model.Todo
	result := database.DB().Table("todo").Where("user_id=? AND time=?", userId, date).Find(&todo)
	if result.Error != nil || result.RowsAffected < 1 {
		return nil, result.Error
	}
	//把待办的日期序列化格式改成YYYY-MM-DD
	for i, v := range todo {
		todo[i].Time = v.Time[:10]
	}
	return todo, result.Error
}

// ListHeaderFromTime 统计，先找到这一段时间有多少个不同的待办表
func ListHeaderFromTime(userId uint, from string) ([]string, error) {
	var headers []string
	now := time.Now().Format("2006-01-02")
	result := database.DB().Table("todo").Distinct("header").Where("user_id=? AND time BETWEEN ? AND ?", userId, from, now).Select("header").Find(&headers)
	return headers, result.Error
}

// ListItemFromTime 通过header寻找这一段时间，这个header下的待办事项
func ListItemFromTime(header string, userId uint, from string) ([]string, error) {
	var items []string
	now := time.Now().Format("2006-01-02")
	result := database.DB().Table("todo").Where("user_id=? AND header=? AND time BETWEEN ? AND ?", userId, header, from, now).Select("todo_item").Find(&items)
	return items, result.Error
}

func DeleteTodo(todoId uint) error {
	result := database.DB().Table("todo").Where("id=?", todoId).Delete(&model.Todo{})
	return result.Error
}
