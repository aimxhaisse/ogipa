package main

// Skeleton to build HTTP/JSON APIs

import (
	"github.com/gin-gonic/gin"
	"log"
)

// Main structure of the API
type PathwarAPI struct {
	r *gin.Engine // Gin's engine
}

// ReturnCode of an API request
type ReturnCode int

const (
	RC_OK             ReturnCode = iota // OK
	RC_INTERNAL_ERROR                   // KO
)

// Response of our API
type Response struct {
	RC ReturnCode
}

// Common stuff

func (p *PathwarAPI) run() {
	p.r.Run(":8080")
}

func (p *PathwarAPI) init() error {
	users = make([]UserJSON, 0)

	users = append(users, UserJSON{"m1ch3l"})

	// users API
	users := p.r.Group("/users")
	{
		// tout ce qui est /users
		// donc là /users/list
		users.GET("/list", func(c *gin.Context) {
			go p.usersList(c)
		})
	}

	// add your own resource

	return nil
}

func NewPathwarAPI() *PathwarAPI {
	return &PathwarAPI{
		gin.Default(),
	}
}

func main() {
	app := NewPathwarAPI()
	err := app.init()
	if err != nil {
		log.Fatalf("failed to init PathwarAPI: %v", err)
	}
	app.run()
}
