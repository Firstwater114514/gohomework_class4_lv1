package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/model"
	"log"
)

var db *sql.DB

func AddUser(username, password, checkQuestion, checkAnswer string) {
	us := model.Users{
		Username:      username,
		Password:      password,
		CheckQuestion: checkQuestion,
		CheckAnswer:   checkAnswer,
	}
	sqlStr := "insert into user(username,password,checkQuestion,checkAnswer) values (?,?,?,?)"
	_, err := db.Exec(sqlStr, us.Username, us.Password, us.CheckQuestion, us.CheckAnswer)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
}

func SelectUser(username string) bool {
	id := 1
	sqlStr := "select username from user where id >= ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var us model.Users
		err := rows.Scan(&us.Username)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return false
		}
		if username == us.Username {
			return false
		}
	}
	return true
}

func IfLogin() string {
	id := 1
	sqlStr := "select login from user where id >= ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return "return1"
	}
	defer rows.Close()
	for rows.Next() {
		var us model.Users
		err := rows.Scan(&us.Login)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return "return2"
		}
		if us.Login == "yes" {
			return "yes"
		}
	}
	return "no"
}

func SelectPasswordFromUsername(username string) string {
	sqlStr := "select password from user where username=?"
	var us model.Users
	err := db.QueryRow(sqlStr, username).Scan(&us.Password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return us.Password
}
func Login(username string) {
	var us model.Users
	us.Login = "yes"
	us.Username = username
	sqlStr := "update user set login=? where username=?"
	_, err := db.Exec(sqlStr, us.Login, us.Username)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func WhoLogin() string {
	login := "yes"
	sqlStr := "select username from user where login=?"
	var us model.Users
	err := db.QueryRow(sqlStr, login).Scan(&us.Username)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return "return"
	}
	return us.Username
}
func ChangePassword(username, newPassword string) {
	var us model.Users
	us.Password = newPassword
	us.Username = username
	sqlStr := "update user set password=? where username=?"
	_, err := db.Exec(sqlStr, us.Password, us.Username)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func CheckQuestion(username string) string {
	sqlStr := "select checkQuestion from user where username=?"
	var us model.Users
	err := db.QueryRow(sqlStr, username).Scan(us.CheckQuestion)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return us.CheckQuestion
}
func CheckAnswer(answer, username string) bool {
	sqlStr := "select checkAnswer from user where username=?"
	var us model.Users
	err := db.QueryRow(sqlStr, username).Scan(&us.CheckAnswer)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	if answer != us.CheckAnswer {
		return false
	}
	return true
}
func IfChange(username string) {
	var us model.Users
	us.Username = username
	us.Change = 'y'
	sqlStr := "update user set change=? where username=?"
	_, err := db.Exec(sqlStr, us.Change, us.Username)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func WhoChang() string {
	change := 'y'
	sqlStr := "select username from user where change=?"
	var us model.Users
	err := db.QueryRow(sqlStr, change).Scan(&us.Username)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return us.Username
}
func ChangeOver() {
	var us1, us2 model.Users
	us2.Change = 'n'
	us1.Change = 'y'
	sqlStr := "update user set change=? where change=?"
	_, err := db.Exec(sqlStr, us1.Change, us2.Change)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func AddComment(comment string) {
	sqlStrF := "select mount from floor where id=?"
	var fl model.Floor
	errF := db.QueryRow(sqlStrF, 1).Scan(&fl.Mount)
	if errF != nil {
		fmt.Printf("scan failed, err:%v\n", errF)
		return
	}
	var co model.Comment
	co.EverComment = comment
	co.Floor = fl.Mount + 1
	sqlStrC := "insert into comments(everyComment,floor) values (?,?)"
	_, errC := db.Exec(sqlStrC, co.EverComment, co.Floor)
	if errC != nil {
		fmt.Printf("insert failed, err:%v\n", errC)
		return
	}
	sqlStr := "update floor set mount=? where id=?"
	_, err := db.Exec(sqlStr, co.Floor, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}

func ReadComment(floor int) string {
	sqlStr := "select everyComment from comments where floor=?"
	var co model.Comment
	err := db.QueryRow(sqlStr, floor).Scan(&co.EverComment)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return ""
	}
	return co.EverComment
}
func FloorMount() int {
	sqlStr := "select mount from floor where id=?"
	var fl model.Floor
	err := db.QueryRow(sqlStr, 1).Scan(&fl.Mount)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return -1
	}
	return fl.Mount
}

func DeleteComment(numI int) {
	sqlStrD := "delete from comments where floor=?"
	_, errD := db.Exec(sqlStrD, numI)
	if errD != nil {
		fmt.Printf("delete failed, err:%v\n", errD)
		return
	}
	sqlStr1 := "select mount from floor where id=?"
	var fl model.Floor
	err1 := db.QueryRow(sqlStr1, 1).Scan(&fl.Mount)
	if err1 != nil {
		fmt.Printf("scan failed, err:%v\n", err1)
		return
	}
	sqlStr2 := "update floor set mount=? where id=?"
	_, err2 := db.Exec(sqlStr2, fl.Mount-1, 1)
	if err2 != nil {
		fmt.Printf("update failed, err:%v\n", err2)
		return
	}
	for floor := numI + 1; ; floor++ {
		sqlStr := "select everyComment from comments where floor=?"
		var co model.Comment
		err := db.QueryRow(sqlStr, floor).Scan(&co.EverComment)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if co.EverComment == "" {
			return
		}
		sqlStrF := "update comments set floor=? where floor=?"
		_, errF := db.Exec(sqlStrF, floor-1, floor)
		if errF != nil {
			fmt.Printf("update failed, err:%v\n", errF)
			return
		}
	}
}

func ClearComments() {
	sqlStr := "update floor set mount=? where id=?"
	_, err := db.Exec(sqlStr, 0, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	for floor := 1; ; floor++ {
		sqlStrC := "select everyComment from comments where floor=?"
		var co model.Comment
		errC := db.QueryRow(sqlStrC, floor).Scan(&co.EverComment)
		if errC != nil {
			fmt.Printf("scan failed, err:%v\n", errC)
			return
		}
		if co.EverComment == "" {
			return
		}
		sqlStrD := "delete from comments where floor=?"
		_, errD := db.Exec(sqlStrD, floor)
		if errD != nil {
			fmt.Printf("delete failed, err:%v\n", errD)
			return
		}
	}
}
func Quit() {
	var us1, us2 model.Users
	us2.Login = "yes"
	us1.Login = "no"
	sqlStr := "update user set login=? where login=?"
	_, err := db.Exec(sqlStr, us1.Login, us2.Login)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
}
func Unsubscribe(username string) {
	sqlStr := "delete from user where username=?"
	_, err := db.Exec(sqlStr, username)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
}
func ClearUsers() {
	sqlStr := "delete from user where login=?"
	_, err := db.Exec(sqlStr, "no")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
}
func InitDB() {
	var err error
	dsn := "root:sqy040213@tcp(127.0.0.1:3306)/comment"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}
