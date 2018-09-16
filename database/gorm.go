package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jaylevin/TMN-API/errs"
)

func GetRecord(db *gorm.DB, t interface{}, conditions ...interface{}) (interface{}, error) {
	result := db.First(t, conditions)
	if result.Error != nil {
		return nil, errs.New(404, result.Error.Error())
	}
	return t, nil
}