package log

import (
	"fmt"
	"github.com/Thenecromance/OurStories/utility/helper"
	"os"
)

func init() {
	if !helper.DirExists(logDir) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	err := helper.CreateIfNotExist(logFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = helper.CreateIfNotExist(errFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	Instance = New()
}
