package main

import (
	"github.com/gin-gonic/gin"
)

// Users API

type UserJSON struct {
	Login string `json:"user" binding:"required"`
}

var users []UserJSON

func (p *PathwarAPI) usersList(c *gin.Context) error {
	type ResponseListUsers struct {
		Response            // stuff common to all responses
		Users    []UserJSON // list of users
	}

	rsp := ResponseListUsers{
		Response{RC_OK},
		users,
	}

	c.JSON(200, rsp)

	return nil
}
