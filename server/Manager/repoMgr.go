package Manager

import (
	"github.com/Thenecromance/OurStories/SQL/MySQL"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
)

type RepoMgr struct {
	repos []Interface.Repository
}

func (r *RepoMgr) Initialize() {

	MySQL.RunScriptFolder("./scripts/MySQL/Initializer")

	log.Info("Initializing Repositories")
	// bind tables to the structures
	for _, repo := range r.repos {
		repo.BindTable()
	}

	// all these stuff has been processed by the scripts
	/*err := MySQL.Default().CreateTablesIfNotExists()
	if err != nil {
		log.Error("Failed to create tables: ", err.Error())
		return
	}*/
	log.Info("Tables created successfully")
}

func (r *RepoMgr) RegisterRepository(repos ...Interface.Repository) {
	for _, repo := range repos {
		r.repos = append(r.repos, repo)
	}
}

func NewRepositoryManager() *RepoMgr {
	return &RepoMgr{}
}
