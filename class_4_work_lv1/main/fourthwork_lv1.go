package main

import (
	_ "github.com/go-sql-driver/mysql"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/api"
	"lanshan_homework/go1.19.2/go_homework/class_4_work_lv1/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
