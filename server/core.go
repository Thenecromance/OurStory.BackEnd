package server

import Interface "github.com/Thenecromance/OurStories/interface"

type core struct {
	option      ServerOption
	controllers []Interface.Controller
}

func (c *core) useTLS() bool {
	return c.option.CertFile != "" && c.option.KeyFile != ""
}
