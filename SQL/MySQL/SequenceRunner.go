package MySQL

import (
	"github.com/Thenecromance/OurStories/utility/File"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
	"strings"
)

type SequenceRunner struct {
	db      *gorp.DbMap
	folder  string
	scripts []string // all the scripts will be splited and stored here
}

func (runner *SequenceRunner) Load() {
	if !File.DirExists(runner.folder) {
		return
	}

	var files []string
	if File.Exists(runner.folder + "/index.idx") {
		files = runner.parse(runner.folder + "/index.idx")
	} else {
		var err error
		files, err = File.ListFiles(runner.folder)
		if err != nil {
			return
		}
	}

	for _, file := range files {
		buf, err := File.ReadFrom(runner.folder + "/" + file)
		if err != nil {
			panic(err)
		}
		for _, sc := range strings.Split(string(buf), ";") {
			if strings.TrimSpace(sc) == "" {
				continue
			}
			runner.scripts = append(runner.scripts, sc)
		}
	}
}

func (runner *SequenceRunner) parse(script string) []string {
	buf, err := File.ReadFrom(script)
	if err != nil {
		panic(err)
	}

	cols := strings.Split(string(buf), "\n")
	var result []string
	for _, col := range cols {
		// remove // comments and empty lines
		if strings.HasPrefix(col, "//") || strings.TrimSpace(col) == "" {
			continue
		}
		// remove trailing comments
		col = strings.Split(col, "//")[0]

		result = append(result, col)
	}
	return result
}

// process the scripts in the folder
func (runner *SequenceRunner) Run() {
	runner.Load()

	for _, script := range runner.scripts {
		if script == "" {
			continue
		}
		_, err := runner.db.Exec(script + ";")
		if err != nil {
			log.Error(err)
			return
		}

	}

}

func NewSequenceRunner(db *gorp.DbMap, folder string) *SequenceRunner {
	return &SequenceRunner{db: db,
		folder: folder}
}

func RunScriptFolder(folder string) {
	runner := NewSequenceRunner(Default(), folder)
	runner.Run()
}
