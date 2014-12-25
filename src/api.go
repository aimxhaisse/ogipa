package main

// - Basic Types

type User struct {
	Login string
}

// - API sugar

type ReturnCode int

const (
	RC_OK             ReturnCode = iota // OK
	RC_INTERNAL_ERROR                   // KO
)

type Response struct {
	RC ReturnCode
}

type ResponseUsersList struct {
	Response
	Users []User
}
