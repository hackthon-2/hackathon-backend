package service

import (
	"encoding/json"
	"errors"
	database "hackthon/database/mysql"
	database2 "hackthon/database/redis"
	"hackthon/model"
	"hackthon/util"
	"strings"
)

var (
	CreateTodoError  = errors.New("creating todo error")
	UpdateTodoError  = errors.New("updating todo error")
	GetTodoListError = errors.New("getting todo error")
	DeleteTodoError  = errors.New("deleting todo error")
)

func CreateTodo(userId uint, input *model.TodoInput) error {
	var todo model.Todo
	util.StructAssign(&todo, input)
	todo.UserID = userId
	if len(input.ToDoItems) == 0 {
		todo.TodoItem = ""
	} else {
		var vals = make([]string, len(input.ToDoItems))
		//获取到的是在结构体里面，所以把每一项的都给序列化
		for i, v := range input.ToDoItems {
			val, _ := json.Marshal(v)
			vals[i] = string(val)
		}
		//然后拼接成字符串赋值给todo
		todo.TodoItem = strings.Join(vals, "/")
	}
	err, row := database.CreateToDo(&todo)
	if err != nil {
		return err
	}
	if row != 1 {
		return CreateTodoError
	}
	data, err := database.ListTodoByTime(userId, input.Time)
	//先找到当天的todo的所有数据
	if err != nil {
		return err
	}
	var response = make([]model.TodoResponse, len(data))
	//根据数据构造结构体，大小就是len(data)
	for i, v := range data {
		//遍历data,先把常项给通过反射赋值
		util.StructAssign(&(response[i]), &v)
		//把数据库里存的这个，看看是不是空的，空的直接赋值一个空的切片
		if v.TodoItem == "" {
			response[i].ToDoItems = []model.ToDoItem{}
		} else {
			//不然就把它给切片了
			val := strings.Split(v.TodoItem, "/")
			response[i].ToDoItems = make([]model.ToDoItem, len(val))
			//然后把每一个切片的json字符串反序列化成ToDoItem
			for index, value := range val {
				var item model.ToDoItem
				_ = json.Unmarshal([]byte(value), &item)
				(response[i].ToDoItems)[index] = item
			}
		}
	}
	dat, _ := json.Marshal(&response)
	err = database2.CreateTodoCache(userId, string(dat), input.Time)
	return err
}

// UpdateTodo 逻辑复用createTodo
func UpdateTodo(userId, todoId uint, input *model.TodoInput) error {
	var todo model.Todo
	util.StructAssign(&todo, input)
	todo.UserID = userId
	if len(input.ToDoItems) == 0 {
		todo.TodoItem = ""
	} else {
		var vals = make([]string, len(input.ToDoItems))
		//获取到的是在结构体里面，所以把每一项的都给序列化
		for i, v := range input.ToDoItems {
			val, _ := json.Marshal(v)
			vals[i] = string(val)
		}
		//然后拼接成字符串赋值给todo
		todo.TodoItem = strings.Join(vals, "/")
	}
	err, row := database.UpdateToDo(todoId, &todo)
	if err != nil {
		return err
	}
	if row != 1 {
		return UpdateTodoError
	}
	data, err := database.ListTodoByTime(userId, input.Time)
	//先找到当天的todo的所有数据
	if err != nil {
		return err
	}
	var response = make([]model.TodoResponse, len(data))
	//根据数据构造结构体，大小就是len(data)
	for i, v := range data {
		//遍历data,先把常项给通过反射赋值
		util.StructAssign(&(response[i]), &v)
		//把数据库里存的这个，看看是不是空的，空的直接赋值一个空的切片
		if v.TodoItem == "" {
			response[i].ToDoItems = []model.ToDoItem{}
		} else {
			//不然就把它给切片了
			val := strings.Split(v.TodoItem, "/")
			response[i].ToDoItems = make([]model.ToDoItem, len(val))
			//然后把每一个切片的json字符串反序列化成ToDoItem
			for index, value := range val {
				var item model.ToDoItem
				_ = json.Unmarshal([]byte(value), &item)
				(response[i].ToDoItems)[index] = item
			}
		}
	}
	dat, _ := json.Marshal(&response)
	err = database2.CreateTodoCache(userId, string(dat), input.Time)
	return err
}

func ListTodo(userID uint, date string) ([]model.TodoResponse, error) {
	data, err := database2.FindTodoCache(userID, date)
	if err != nil || data == "" {
		dat, e := database.ListTodoByTime(userID, date)
		//先找到当天的todo的所有数据
		if e != nil {
			return []model.TodoResponse{}, GetTodoListError
		}
		if len(dat) < 1 {
			return []model.TodoResponse{}, nil
		}
		var response = make([]model.TodoResponse, len(data))
		//根据数据构造结构体，大小就是len(data)
		for i, v := range dat {
			//遍历data,先把常项给通过反射赋值
			util.StructAssign(&(response[i]), &v)
			//把数据库里存的这个，看看是不是空的，空的直接赋值一个空的切片
			if v.TodoItem == "" {
				response[i].ToDoItems = []model.ToDoItem{}
			} else {
				//不然就把它给切片了
				val := strings.Split(v.TodoItem, "/")
				response[i].ToDoItems = make([]model.ToDoItem, len(val))
				//然后把每一个切片的json字符串反序列化成ToDoItem
				for index, value := range val {
					var item model.ToDoItem
					_ = json.Unmarshal([]byte(value), &item)
					(response[i].ToDoItems)[index] = item
				}
			}
		}
		d, _ := json.Marshal(&response)
		err = database2.CreateTodoCache(userID, string(d), date)
		return response, err
	}
	if len(data) < 1 {
		err = database2.DeleteTodoFieldCache(userID, date)
		return nil, err
	}
	var res []model.TodoResponse
	_ = json.Unmarshal([]byte(data), &res)
	return res, nil
}

func DeleteTodo(userId, todoID uint) error {
	todo, row, err := database.FindToDoById(todoID)
	if row != 1 {
		return DeleteTodoError
	}
	if err != nil {
		return err
	}
	err = database.DeleteTodo(todoID)
	if err != nil {
		return err
	}
	data, err := database.ListTodoByTime(userId, todo.Time)
	if err != nil {
		return err
	}
	var response = make([]model.TodoResponse, len(data))
	//根据数据构造结构体，大小就是len(data)
	for i, v := range data {
		//遍历data,先把常项给通过反射赋值
		util.StructAssign(&(response[i]), &v)
		//把数据库里存的这个，看看是不是空的，空的直接赋值一个空的切片
		if v.TodoItem == "" {
			response[i].ToDoItems = []model.ToDoItem{}
		} else {
			//不然就把它给切片了
			val := strings.Split(v.TodoItem, "/")
			response[i].ToDoItems = make([]model.ToDoItem, len(val))
			//然后把每一个切片的json字符串反序列化成ToDoItem
			for index, value := range val {
				var item model.ToDoItem
				_ = json.Unmarshal([]byte(value), &item)
				(response[i].ToDoItems)[index] = item
			}
		}
	}
	if len(data) < 1 {
		return database2.DeleteTodoFieldCache(userId, todo.Time)
	}
	d, _ := json.Marshal(&response)
	return database2.CreateTodoCache(userId, string(d), todo.Time)
}
