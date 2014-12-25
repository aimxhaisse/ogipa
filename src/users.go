package main

import (
	"github.com/gin-gonic/gin"
	"github.com/dancannon/gorethink"
)

// Users API

func (p *PathwarAPI) usersNew(c *gin.Context) error {
	var user User
	if ! c.Bind(&user) {
		c.JSON(400, Response{RC_INVALID_PARAMS})
		return nil
	}
	_, err := gorethink.Table("users").Insert(user).RunWrite(p.db)
	return err
}

func (p *PathwarAPI) usersList(c *gin.Context) error {
	users := make([]User, 0)
	users = append(users, User{"m1ch3l"})

	rsp := ResponseUsersList{
		Response{RC_OK},
		users,
	}

	c.JSON(200, rsp)

	return nil
}
