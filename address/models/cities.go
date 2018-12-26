package models

import (
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"upper.io/db.v3/lib/sqlbuilder"
)

type City struct {
	CityId      int     `db:"city_id" json:"city_id"`
	Name        string  `db:"name" json:"name"`
	Citycode    string  `db:"citycode" json:"citycode"`
	ProvinceId  int     `db:"province_id" json:"province_id"`
	Longtitude  float64 `db:"center_longtitude" json:"longtitude"`
	Latitude    float64 `db:"center_latitude" json:"latitude"`
	WeightValue int     `db:"weight_value" json:"weight_value"`
}

func (City) TableName() string {
	return "cities"
}

func GetCitiesByProvinceId(id int, db sqlbuilder.Database) (cities []City, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetCitiesByProvinceId] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	city := db.Collection(City{}.TableName())
	cities = []City{}
	if err = city.Find("province_id", id).All(&cities); err != nil {
		log.ErrorRaw("[GetCitiesByProvinceId] read `cities` records failed. err=%s", err)
		return
	}
	return
}

func GetCityById(id int, db sqlbuilder.Database) (city *City, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetCityById] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	city = new(City)
	if err = db.Collection(city.TableName()).Find("city_id", id).One(city); err != nil {
		log.ErrorRaw("[GetCityById] read `city` record failed. err=%s", err)
		return
	}
	return
}
