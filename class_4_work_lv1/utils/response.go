package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}
func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  500,
		"message": message,
	})
}
func Question(c *gin.Context, message, tip string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip":     tip,
	})
}
func AnswerRight(c *gin.Context, message, tip1 string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip1":    tip1,
	})
}
func CommentsWall(c *gin.Context, comments string) {
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"comments": comments,
	})
}
func Comment(c *gin.Context, floor int, comment string) {
	c.JSON(http.StatusOK, gin.H{
		"floor":   floor,
		"comment": comment,
	})
}
func LoginSuccess(c *gin.Context, message, tip1, tip2, tip3, tip4, tip5 string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
		"tip1":    tip1,
		"tip2":    tip2,
		"tip3":    tip3,
		"tip4":    tip4,
		"tip5":    tip5,
	})
}
