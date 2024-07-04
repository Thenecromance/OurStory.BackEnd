package repository

import (
	SQL "github.com/Thenecromance/OurStories/SQL/redis"
	"github.com/go-redis/redis/v8"
	"gopkg.in/gorp.v2"
)

type RouteControlRepository interface {
}

type routeControlRepository struct {
	db  *gorp.DbMap
	cli *redis.Client
}

func (r *routeControlRepository) initialize() {
	if r.db == nil {
		return
	}

}
func (r *routeControlRepository) SQLToRedisSync() {
	if r.db == nil || r.cli == nil {
		return
	}

	r.db.Exec("SELECT * FROM route_control")
}

func (r *routeControlRepository) RedisToSQLSync() {
	if r.db == nil || r.cli == nil {
		return
	}
}

func NewRouteControl(db *gorp.DbMap) RouteControlRepository {
	opt := SQL.LoadSetting().ToRedisOption()

	return &routeControlRepository{
		db:  db,
		cli: redis.NewClient(opt),
	}
}
