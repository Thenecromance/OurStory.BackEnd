//go:build linux

package SQL

import (
	"github.com/Thenecromance/OurStories/utility/log"
	"os/exec"
	"strings"
)

func init() {
	cfg := defaultConfig()
	if cfg.Host == "" ||
		cfg.Host == "127.0.0.1" ||
		cfg.Host == "localhost" {
		if !runningOnPlatform() {
			startSQL()
		}
	}
}

func runningOnPlatform() bool {
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

func startSQL() {
	cmd := exec.Command("service", "mysql", "start")
	err := cmd.Run()
	if err != nil {
		//log.Errorf("MakeAsTree failed:%s", err)
	}
}

func stopSQL() {
	cmd := exec.Command("service", "mysql", "stop")
	err := cmd.Run()
	if err != nil {
		log.Errorf("MakeAsTree failed:%s", err)
	}
}
