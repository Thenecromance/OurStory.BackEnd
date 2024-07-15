package main

import (
	"fmt"
	"github.com/Thenecromance/OurStories/SQL/MySQL"
	"github.com/Thenecromance/OurStories/application/models"
	"time"
)

func main() {
	MySQL.RunScriptFolder("../../scripts/MySQL/Initializer")

	usr := models.User{}
	MySQL.Default().AddTableWithName(usr, "Users")

	usr.UserName = "test"
	usr.NickName = "test"
	usr.Password = "test"
	usr.Email = "123"
	usr.Salt = "123"
	usr.UserId = 1
	usr.Birthday = time.Now().AddDate(-10, 0, 0)
	usr.CreatedAt = time.Now().AddDate(-10, 0, 0)
	usr.LastLogin = time.Now().AddDate(-10, 0, 0)

	/*	err := MySQL.Default().Insert(&usr)
		if err != nil {
			fmt.Println(err)
			return
		}*/

	MySQL.Default().Get(&usr, "select * from Users where user_id = ?", 1)

	fmt.Println(usr.Birthday, "\n", usr.Birthday.UnixMilli())

}
