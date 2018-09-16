package database

import (
	"github.com/jaylevin/TMN-API/errs"
	"github.com/jinzhu/gorm"
)

// Just a wrapper function for dynamic error handling
func GetRecord(db *gorm.DB, t interface{}, conditions ...interface{}) (interface{}, error) {
	result := db.First(t, conditions)
	if result.Error != nil {
		return nil, errs.New(404, result.Error.Error())
	}
	return t, nil
}
