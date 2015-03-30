package service

import (
	"github.com/AntonMaltsev/uis_dockable/api"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

type TodoResource struct {
	db gorm.DB
}

func (tr *TodoResource) CreateUser(c *gin.Context) {
	var todo api.Todo

	if !c.Bind(&todo) {
		c.JSON(400, api.NewError("problem decoding body"))
		return
	}
	todo.Status = api.TodoStatus
	todo.Created = int32(time.Now().Unix())

	

	tr.db.Save(&todo)

	c.JSON(201, todo)
}

func (tr *TodoResource) GetAllUsers(c *gin.Context) {
	var todos []api.Todo

	tr.db.Order("created desc").Find(&todos)

	c.JSON(200, todos)
}

func (tr *TodoResource) GetUser(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding id sent"))
		return
	}

	var todo api.Todo

	if tr.db.First(&todo, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, todo)
	}
}

func (tr *TodoResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
