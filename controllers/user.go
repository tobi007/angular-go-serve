package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tobi007/angular-go-serve/models"
	"github.com/tobi007/angular-go-serve/repository"
	"github.com/tobi007/angular-go-serve/repository/dao"
	"github.com/tobi007/angular-go-serve/util"
	"io"
	"io/ioutil"
	"net/http"
)

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		dao: dao.NewuserDao(db),
	}
}

type UserController struct {
	dao repository.UserRepo
}

// AddTodoHandler adds a new todo to the todo list
func (u UserController) Create(c *gin.Context) {
	user, statusCode, err := convertHTTPBodyToUser(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	user.Password, err = util.GenerateHash(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error")
		return
	}
	_, err = u.dao.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"User Created Successfully", "data": user.Email})
}

func convertHTTPBodyToUser(httpBody io.ReadCloser) (models.User, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToUser(body)
}


func convertJSONBodyToUser(jsonBody []byte) (models.User, int, error) {
	var todoItem models.User
	err := json.Unmarshal(jsonBody, &todoItem)
	if err != nil {
		return models.User{}, http.StatusBadRequest, err
	}
	return todoItem, http.StatusOK, nil
}
