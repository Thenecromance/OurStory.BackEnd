package SQL

import "github.com/Thenecromance/OurStories/utility/log"

type gLogger struct {
}

func (g *gLogger) Printf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}
