package orm

import (
	"github.com/dsogari/barber-shop/graph/generated"
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
	Db.AutoMigrate(&generated.Shop{})
	Db.AutoMigrate(&generated.Barber{})
	Db.AutoMigrate(&generated.Service{})
	Db.AutoMigrate(&generated.Client{})
	Db.AutoMigrate(&generated.Attendance{})
}
