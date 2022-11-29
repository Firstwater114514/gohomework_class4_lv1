package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/api/middleware"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/dao"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/model"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/utils"
	"strconv"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.Register{}); err != nil {
		utils.RespSuccess(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后登录")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "用户名已被使用")
		return
	}
	password := c.PostForm("password")
	checkQuestion := c.PostForm("check question")
	checkAnswer := c.PostForm("check answer")
	dao.AddUser(username, password, checkQuestion, checkAnswer)
	utils.RespSuccess(c, "注册成功!快去登录吧~")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.Login{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "yes" {
		utils.RespFail(c, "已有账号登录，请退出在线账号后登录")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if flag {
		utils.RespFail(c, "该用户不存在")
		return
	}
	selectPassword := dao.SelectPasswordFromUsername(username)
	if selectPassword != password {
		utils.RespFail(c, "密码错误！")
		return
	} else if selectPassword == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	dao.Login(username)
	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "sqy",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.LoginSuccess(c, tokenString, "将请求中的login改成add comment并且输入comment就可以写留言啦", "将请求中的login改成delete comment并且输入想删除的留言序号num", "将请求中的login改成change password并且输入new password可以更改密码哦", "注销账号请使用unsubscribe", "记得关闭窗口前一定要退出账号哦!!!!!用quit")
}

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func changePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Change{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	username := dao.WhoLogin()
	if username == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	password := dao.SelectPasswordFromUsername(username)
	newPassword := c.PostForm("new password")
	if password == newPassword {
		utils.RespFail(c, "新密码与旧密码相同~")
		return
	} else if password == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	dao.ChangePassword(username, newPassword)
	dao.Quit()
	utils.RespSuccess(c, "更改密码成功，请重新登录~")
	return
}
func forgetPassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Forget{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	username := c.PostForm("username")
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "该用户不存在")
		return
	}
	if dao.CheckQuestion(username) == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	utils.Question(c, dao.CheckQuestion(username), "把请求中的forget password改成answer并且输入answer、username")
}
func answer(c *gin.Context) {
	if err := c.ShouldBind(&model.Question{}); err != nil {
		utils.RespFail(c, "部分数据未输入，请检查")
		return
	}
	username := c.PostForm("username")
	answer := c.PostForm("answer")
	check := dao.CheckAnswer(answer, username)
	if !check {
		utils.RespFail(c, "验证答案不正确，请重新输入")
		return
	}
	dao.IfChange(username)
	utils.AnswerRight(c, "验证成功！", "将请求中的answer改成update password并且输入new password")
}
func updatePassword(c *gin.Context) {
	if err := c.ShouldBind(&model.Change{}); err != nil {
		utils.RespFail(c, "请输入新的密码")
		return
	}
	username := dao.WhoChang()
	if username == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	newPassword := c.PostForm("new password")
	dao.ChangePassword(username, newPassword)
	dao.ChangeOver()
	utils.RespSuccess(c, "更改密码成功！请重新登录哦")
	return
}
func addComment(c *gin.Context) {
	if err := c.ShouldBind(&model.AddComment{}); err != nil {
		utils.RespFail(c, "请输入你的留言")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	username := dao.WhoLogin()
	if username == "" {
		utils.RespFail(c, "scan failed")
		return
	}
	comment := c.PostForm("comment")
	realComment := username + ":" + comment
	dao.AddComment(realComment)
	utils.RespSuccess(c, "留言成功，把请求中的add comment改成scan comments查看留言板")
}
func scanComments(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	if dao.FloorMount() == 0 {
		utils.RespSuccess(c, "暂时还没有留言啦，第一个留言就交给你咯~")
		return
	}
	utils.CommentsWall(c, "留言板：")
	for floor := 1; ; floor++ {
		rc := dao.ReadComment(floor)
		if rc == "" {
			return
		}
		utils.Comment(c, floor, rc)
	}
}
func deleteComment(c *gin.Context) {
	if err := c.ShouldBind(&model.DeleteComment{}); err != nil {
		utils.RespFail(c, "请输入想删除的留言序号哦")
		return
	}
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	num := c.PostForm("num")
	fm := dao.FloorMount()
	if fm == 0 {
		utils.RespSuccess(c, "还没有人留言哦，没有留言可以删呢~")
		return
	} else if fm == -1 {
		utils.RespFail(c, "scan failed")
	}
	numI, errN := strconv.Atoi(num)
	if errN != nil {
		panic(errN)
	}
	if numI > fm {
		utils.RespFail(c, "没有该序号的留言")
		return
	}
	dao.DeleteComment(numI)
	if dao.FloorMount() == 0 {
		utils.RespSuccess(c, "成功删除留言，留言板暂无留言")
		return
	}
	utils.CommentsWall(c, "留言板：")
	for floor := 1; ; floor++ {
		rc := dao.ReadComment(floor)
		if rc == "" {
			return
		}
		utils.Comment(c, floor, rc)
	}
}
func clearComments(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	if dao.FloorMount() == 0 {
		utils.RespFail(c, "留言板本来就是空的哦~")
		return
	}
	dao.ClearComments()
	utils.RespSuccess(c, "清除留言板成功")
}
func quit(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "没有登录的账户")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	dao.Quit()
	utils.RespSuccess(c, "成功退出账号")
	return
}
func unsubscribe(c *gin.Context) {
	if dao.IfLogin() == "no" {
		utils.RespFail(c, "请先登录再进行其他操作")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	username := dao.WhoLogin()
	dao.Unsubscribe(username)
	utils.RespSuccess(c, "注销账户成功")
}
func clearAll(c *gin.Context) {
	if dao.IfLogin() == "yes" {
		utils.RespFail(c, "该功能仅退出账号后可使用")
		return
	} else if dao.IfLogin() == "return1" {
		utils.RespFail(c, "query failed")
		return
	} else if dao.IfLogin() == "return2" {
		utils.RespFail(c, "scan failed")
		return
	}
	dao.ClearComments()
	dao.ClearUsers()
	utils.RespSuccess(c, "成功初始化该系统！")
}
