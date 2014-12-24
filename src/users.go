package main

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Users API

type UserJSON struct {
	Login string `json:"user" binding:"required"`
}

func (p *PathwarAPI) usersNew(c *gin.Context) error {
	return errors.New("Not implemented")
}

func (p *PathwarAPI) usersList(c *gin.Context) error {
	type ResponseListUsers struct {
		Response            // stuff common to all responses
		Users    []UserJSON // list of users
	}

	users := make([]UserJSON, 0)
	users = append(users, UserJSON{"m1ch3l"})

	rsp := ResponseListUsers{
		Response{RC_OK},
		users,
	}

	c.JSON(200, rsp)

	return nil
}
