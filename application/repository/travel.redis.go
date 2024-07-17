package repository

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/cache/redisCache"
	"github.com/Thenecromance/OurStories/utility/log"
	"strconv"
	"time"
)

var (
	cacheTravelExpireTime = 3600 * time.Second
)

const (
	prefixTravelIdToObject = "travel.id"
)

type travelRedis struct {
	cli Interface.ICache
}

func (t travelRedis) BindTable() error {
	log.Warn("do not use this method with cache, use db instead")
	return nil
}

func (t travelRedis) CreateTravel(info *models.Travel) error {
	t.cli.Prefix(prefixTravelIdToObject)
	return t.cli.Set(strconv.FormatInt(info.Id, 10), info, cacheTravelExpireTime)
}

func (t travelRedis) DeleteTravel(travelId int64) error {
	t.cli.Prefix(prefixTravelIdToObject)
	return t.cli.Delete(strconv.FormatInt(travelId, 10))
}

func (t travelRedis) UpdateTravel(info *models.Travel) error {
	//TODO implement me
	panic("implement me")
}

func (t travelRedis) GetTravelByID(travelId int64) (*models.Travel, error) {
	t.cli.Prefix(prefixTravelIdToObject)
	sId := strconv.FormatInt(travelId, 10)
	obj, err := t.cli.Get(sId)
	if err != nil {
		return nil, err
	}
	m := &models.Travel{}
	err = json.Unmarshal([]byte(obj.(string)), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (t travelRedis) GetTravelByOwner(owner int64) ([]models.Travel, error) {
	//TODO implement me
	panic("implement me")
}

func (t travelRedis) GetTravelByLocation(location string) ([]models.Travel, error) {
	//TODO implement me
	panic("implement me")
}

func (t travelRedis) GetTravelByState(state int) ([]models.Travel, error) {
	//TODO implement me
	panic("implement me")
}

func (t travelRedis) GetTravelListByID(id int64) ([]models.Travel, error) {
	//TODO implement me
	panic("implement me")
}

func newTravelCache() TravelRepository {
	return &travelRedis{
		cli: redisCache.NewCache(),
	}
}
