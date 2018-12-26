package models

import (
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Street struct {
	StreetId    int     `db:"street_id" json:"street_id"`
	Name        string  `db:"name" json:"name"`
	Adcode      string  `db:"adcode" json:"adcode"`
	Citycode    string  `db:"citycode" json:"citycode"`
	ProvinceId  int     `db:"province_id" json:"province_id"`
	CityId      int     `db:"city_id" json:"city_id"`
	DistrictId  int     `db:"district_id" json:"district_id"`
	Longtitude  float64 `db:"center_longtitude" json:"longtitude"`
	Latitude    float64 `db:"center_latitude" json:"latitude"`
	WeightValue int     `db:"weight_value" json:"weight_value"`
}

func (Street) TableName() string {
	return "streets"
}

func GetStreetsByDistrictId(id int, db sqlbuilder.Database) (streets []Street, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetStreetsByDistrictId] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	sess := db.Collection(Street{}.TableName())
	if err = sess.Find("district_id", id).All(&streets); err != nil {
		log.ErrorRaw("[GetStreetsByDistrictId] read `streets` records failed. err=%s", err)
		return
	}
	return
}

func GetStreetById(id int, db sqlbuilder.Database) (s *Street, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetStreetById] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	s = new(Street)
	sess := db.Collection(s.TableName())
	if err = sess.Find("street_id", id).One(s); err != nil {
		log.ErrorRaw("[GetStreetById] read `street` record failed. err=%s", err)
		return
	}
	return
}
