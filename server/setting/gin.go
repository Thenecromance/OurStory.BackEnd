package setting

import Config "github.com/Thenecromance/OurStories/utility/config"

type Gin struct {
	HtmlFiles []string          `ini:"html_files" json:"html_files" yaml:"html_files"`
	NoRoute   string            `ini:"no_route" json:"no_route" yaml:"no_route"`
	NoMethod  string            `ini:"no_method" json:"no_method" yaml:"no_method"`
	ReMap     map[string]string `ini:"re_map" json:"re_map" yaml:"re_map"`
	//StaticResource string            `ini:"static_resource" json:"static_resource" yaml:"static_resource"`
	Redirects map[string]string `ini:"redirects" json:"redirects" yaml:"redirects"`
}

func (c *Gin) load() {
	Config.LoadToObject("gin", c)
}

func NewGinSetting() *Gin {
	cfg := &Gin{
		HtmlFiles: []string{},
		NoRoute:   "",
		NoMethod:  "",
		ReMap:     map[string]string{},
		Redirects: map[string]string{},
	}
	cfg.load()
	return cfg
}
