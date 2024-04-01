package main

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/Dashboard"
	"github.com/Thenecromance/OurStories/backend/Location"
	"github.com/Thenecromance/OurStories/backend/Travel"
	"github.com/Thenecromance/OurStories/backend/User"
	"github.com/Thenecromance/OurStories/backend/Weather"
	"github.com/Thenecromance/OurStories/backend/api"
	"github.com/Thenecromance/OurStories/base/SQL"
	Config "github.com/Thenecromance/OurStories/base/config"
	Interface "github.com/Thenecromance/OurStories/interface"
	"github.com/Thenecromance/OurStories/server"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"time"
)

func loadServerComponent() *server.Server {

	svr := server.New()

	//dashboard controller for control the dashboard text values' change

	svr.LoadComponent(
		api.NewController(),
		Dashboard.NewController(),

		User.NewController(),
		Location.NewController(),
		Weather.NewController(),
		Travel.NewController(),
	)
	return svr
}

func initGinMode() {
	mode := Config.GetBool("Server", "ReleaseMode")
	if mode {
		gin.SetMode(gin.ReleaseMode)
	} else if mode {
		Config.SetBool("Server", "ReleaseMode", false)
	}

}

func nodeTest() {
	temp := Interface.NewNode("credit", "store")
	ctrl := Interface.NewRootNode()
	api := Interface.NewNode("/", "api")
	ctrl.Load(api,
		Interface.NewNode("api", "user"),
		Interface.NewNode("api", "location"),
		Interface.NewNode("api", "weather"),
		Interface.NewNode("api", "travel"),
		Interface.NewNode("api", "credit"),
		Interface.NewNode("/", "dashboard"),
		Interface.NewNode("dashboard", "topcard"),
		Interface.NewNode("dashboard", "sidenavbar"),
		Interface.NewNode("dashboard", "main"),
		Interface.NewNode("credit", "bank"),
		Interface.NewNode("a", "D"), // this node will be ignored
		temp,
	)

	ctrl.MakeAsTree()

	fmt.Println(ctrl.String())
	fmt.Println(temp.Path())
}
func snowflakeTest() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", node.Generate().Int64())
	time.Sleep(1 * time.Second)
	fmt.Printf("ID       : %d\n", node.Generate().Int64())
}
func main() {

	initGinMode()
	svr := loadServerComponent()
	SQL.Initialize()
	svr.Run(":8080")

}
