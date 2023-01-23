package orm

import (
	"github.com/dsogari/barber-shop/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func SetupDatabase(filename string) {
	println("Opening " + filename)
	var err error
	Db, err = gorm.Open(sqlite.Open(filename+"?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&model.Shop{})
	Db.AutoMigrate(&model.Barber{})
	Db.AutoMigrate(&model.Service{})
	Db.AutoMigrate(&model.Client{})
	Db.AutoMigrate(&model.Attendance{})
}
