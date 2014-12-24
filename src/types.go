package main

import (
	"github.com/gin-gonic/gin"
)

// Main structure of the API
type PathwarAPI struct {
	r   *gin.Engine // Gin's engine
	cfg *Config     // App's config
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
