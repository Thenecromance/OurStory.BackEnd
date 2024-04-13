package log

import (
	"fmt"
	"github.com/Thenecromance/OurStories/base/utils"
	"os"
)

func init() {
	if !utils.DirExists(logDir) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	err := utils.CreateIfNotExist(logFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = utils.CreateIfNotExist(errFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	Instance = New()
}
