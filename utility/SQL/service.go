package SQL

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"os/exec"
	"strings"
)

func runningOnOS() bool {
	cmd := exec.Command("service", "mysql", "status")
	output, err := cmd.Output()
	if err != nil {
		//log.Errorf("MakeAsTree failed:%s", err)
		return false
	}
	if strings.Contains(string(output), "Uptime") {
		return true
	}
	return false
}

func start() {
	cmd := exec.Command("service", "mysql", "start")
	err := cmd.Run()
	if err != nil {
		//log.Errorf("MakeAsTree failed:%s", err)
	}
}
func stop() {
	cmd := exec.Command("service", "mysql", "stop")
	err := cmd.Run()
	if err != nil {
		log.Errorf("MakeAsTree failed:%s", err)
	}
}
