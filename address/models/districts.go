package models

import (
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"upper.io/db.v3/lib/sqlbuilder"
)

type District struct {
	DistrictId  int     `db:"district_id" json:"district_id"`
	Name        string  `db:"name" json:"name"`
	Adcode      string  `db:"adcode" json:"adcode"`
	Citycode    string  `db:"citycode" json:"citycode"`
	ProvinceId  int     `db:"province_id" json:"province_id"`
	CityId      int     `db:"city_id" json:"city_id"`
	Longtitude  float64 `db:"center_longtitude" json:"longtitude"`
	Latitude    float64 `db:"center_latitude" json:"latitude"`
	WeightValue int     `db:"weight_value" json:"weight_value"`
}

func (District) TableName() string {
	return "districts"
}

func GetDistrictsByCityId(id int, db sqlbuilder.Database) (districts []District, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetDistrictsById] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	districts = []District{}
	sess := db.Collection(District{}.TableName())
	if err = sess.Find("city_id=?", id).All(&districts); err != nil {
		log.ErrorRaw("[GetDistrictsById] read `districts` records failed. err=%s", err)
		return
	}
	return
}

func GetDistrictById(id int, db sqlbuilder.Database) (district *District, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetDistrictById] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	district = new(District)
	sess := db.Collection(District{}.TableName())
	if err = sess.Find("district_id", id).One(district); err != nil {
		log.ErrorRaw("[GetDistrictById] read `district` record failed. err=%s", err)
		return
	}
	return
}
