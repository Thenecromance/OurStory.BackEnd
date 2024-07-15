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
	MySQL.Default().AddTableWithName(models.User{}, "Users")

	/*	usr.UserName = "test"
		usr.NickName = "test"
		usr.Password = "test"
		usr.Email = "123"
		usr.Salt = "123"
		usr.UserId = 1
		usr.Birthday = time.Now().AddDate(-10, 0, 0).UnixMilli()
		usr.CreatedAt = time.Now().AddDate(-10, 0, 0).UnixMilli()
		usr.LastLogin = time.Now().AddDate(-10, 0, 0).UnixMilli()
		fmt.Println(time.Now().AddDate(-10, 0, 0).UnixMilli())
		err := MySQL.Default().Insert(&usr)
		if err != nil {
			fmt.Println(err)
			return
		}*/

	err := MySQL.Default().SelectOne(&usr, "SELECT * FROM Users WHERE user_id = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(usr.UserName)
	fmt.Println(usr.Birthday, "\n", time.UnixMilli(usr.Birthday))

}
