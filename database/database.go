// database/database.go

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// GlobalDB a global db object will be used across different packages
var GlobalDB *gorm.DB

// InitDatabase creates a sqlite db
func InitDatabase() (err error) {
	GlobalDB, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
