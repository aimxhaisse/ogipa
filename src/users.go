package main

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Users API

func (p *PathwarAPI) usersNew(c *gin.Context) error {
	return errors.New("Not implemented")
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
