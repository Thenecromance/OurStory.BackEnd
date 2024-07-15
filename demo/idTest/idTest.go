package main

import (
	"github.com/Thenecromance/OurStories/SQL/MySQL"
)

func main() {

	MySQL.RunScriptFolder("../../scripts/MySQL/Initializer")

	return
	//creator := Transaction.New()
	//
	//id := creator.NextID(10005596, 100)
	//fmt.Println(id)
	//// format the number in binary
	//id = 1720544286918
	//var binary string
	//for id > 0 {
	//	binary = fmt.Sprintf("%d", id%2) + binary
	//	id = id / 2
	//}
	//fmt.Println(binary)
	/*
		fmt.Println(binary[0:41])
		fmt.Println(binary[41 : 41+12])
		fmt.Println(binary[41+12:])*/

}
