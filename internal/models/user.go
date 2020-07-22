// package name
package models

// lirary imported
import (	
) 

// user structure (document in mongodb)
type User struct {
	Username string             `form:"username" json:"username"`
	Password string             `form:"password" json:"password"`
	Status   string             `form:"status" json:"status"`
}

// login structure 
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type TokenInfoUser struct {
	Status    string
	UserName  string
}