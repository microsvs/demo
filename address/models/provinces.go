package models

import (
	pdb "github.com/microsvs/base/cmd/db"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
	"github.com/microsvs/base/pkg/rpc"
	"upper.io/db.v3/lib/sqlbuilder"
)

type Province struct {
	ProvinceId  int     `db:"province_id" json:"province_id"`
	Name        string  `db:"name" json:"name"`
	Citycode    string  `db:"citycode" json:"citycode"`
	Adcode      string  `db:"adcode" json:"adcode"`
	Longtitude  float64 `db:"center_longtitude" json:"longtitude"`
	Latitude    float64 `db:"center_latitude" json:"latitude"`
	WeightValue int     `db:"weight_value" json:"weight_value"`
}

func (Province) TableName() string {
	return "provinces"
}

func GetProvinces(db sqlbuilder.Database) (ps []Province, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetProvinces] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	sess := db.Collection(Province{}.TableName())
	if err = sess.Find().All(&ps); err != nil {
		log.ErrorRaw("[GetProvinces] read `provinces` records failed.")
		return
	}
	return
}

func GetProvinceById(id int, db sqlbuilder.Database) (ps *Province, err error) {
	if db == nil {
		if db = pdb.SlaveDB(rpc.FGSAddress); db == nil {
			log.ErrorRaw("[GetProvinces] init mysql connection failed.")
			return nil, errors.FGEDBError
		}
	}
	ps = new(Province)
	sess := db.Collection(Province{}.TableName())
	if err = sess.Find("province_id", id).One(ps); err != nil {
		log.ErrorRaw("[GetProvinceById] read `provinces` record failed. err=%s", err)
		return
	}
	return
}
