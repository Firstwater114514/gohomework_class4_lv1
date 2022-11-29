package model

import "github.com/dgrijalva/jwt-go"

type Register struct {
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
	CheckQuestion string `form:"check question" json:"check question" binding:"required"`
	CheckAnswer   string `form:"check answer" json:"check answer" binding:"required"`
}
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type Change struct {
	NewPassword string `form:"new password" json:"new password" binding:"required"`
}
type Forget struct {
	Username string `form:"username" json:"username" binding:"required"`
}
type Question struct {
	Answer   string `form:"answer" json:"answer" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
}
type AddComment struct {
	Comment string `form:"comment" json:"comment" binding:"required"`
}
type DeleteComment struct {
	Num string `form:"num" json:"num" binding:"required"`
}
type Users struct {
	Id            int
	Username      string
	Password      string
	CheckQuestion string
	CheckAnswer   string
	Login         string
	Change        rune
}
type Comment struct {
	Num         int
	EverComment string
	Floor       int
}
type Floor struct {
	Id    int
	Mount int
}
