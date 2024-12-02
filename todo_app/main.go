package main

//cd C:\Users\kanri\go\src\todo_app
//go run main.go
//sqlite3 webapp.sql
//SELECT * FROM todos;
//.exit

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("t-nezu@winas.jp")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)

}
