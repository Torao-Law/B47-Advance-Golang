package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Todos struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

var todos = []Todos{
	{
		Id:     "1",
		Title:  "Front End Engineer",
		IsDone: true,
	},
	{
		Id:     "2",
		Title:  "Mobile Dev Engineer",
		IsDone: false,
	},
	{
		Id:     "3",
		Title:  "Java Developer",
		IsDone: false,
	},
}

func main() {
	e := echo.New()

	// routes
	e.GET("/todos", FindTodos)
	e.GET("/todo/:id", GetTodo)
	e.POST("/todo", CreateTodo)
	e.PATCH("/todo/:id", UpdateTodo)
	e.DELETE("/todo/:id", DeleteTodo)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func FindTodos(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(todos)
}

func GetTodo(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	id := c.Param("id")

	var todoData Todos
	var isGetTodo = false

	for _, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todoData = todo
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().WriteHeader(200)
	return json.NewEncoder(c.Response()).Encode(todoData)
}

func CreateTodo(c echo.Context) error {
	var todo Todos

	json.NewDecoder(c.Request().Body).Decode(&todo)

	todos = append(todos, todo)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(todos)
}

func UpdateTodo(c echo.Context) error {
	var id = c.Param("id")
	var data Todos
	isGetTodo := false

	json.NewDecoder(c.Request().Body).Decode(&data)

	for idx, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			todos[idx] = data
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(data)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	var isGetTodo = false
	var index = 0

	for i, todo := range todos {
		if id == todo.Id {
			isGetTodo = true
			index = i
		}
	}

	if !isGetTodo {
		c.Response().WriteHeader(400)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	todos = append(todos[:index], todos[index+1:]...)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(todos)
}
