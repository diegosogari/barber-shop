package resolvers

import (
	"fmt"
	"strings"

	"github.com/dsogari/barber-shop/graph/generated"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var Db *gorm.DB

func MigrateSchema(db *gorm.DB) {
	Db = db
	Db.AutoMigrate(&generated.Shop{})
	Db.AutoMigrate(&generated.Barber{})
	Db.AutoMigrate(&generated.Service{})
	Db.AutoMigrate(&generated.Client{})
	Db.AutoMigrate(&generated.Attendance{})
}

func getIntArray(columnName string, array []int) string {
	var result string
	if len(array) > 0 {
		str := strings.ReplaceAll(fmt.Sprint(array), " ", ",")
		result = fmt.Sprintf(" AND %s IN (%s)", columnName, str[1:len(str)-1])
	}
	return result
}
