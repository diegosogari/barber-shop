package resolvers

import (
	"fmt"
	"strings"

	"github.com/dsogari/barber-shop/graph/model"
	"gorm.io/gorm"
)

//go:generate gqlgen

type Resolver struct{}

var Db *gorm.DB

func MigrateSchema(db *gorm.DB) {
	Db = db
	Db.AutoMigrate(&model.Shop{})
	Db.AutoMigrate(&model.Barber{})
	Db.AutoMigrate(&model.Service{})
	Db.AutoMigrate(&model.Client{})
	Db.AutoMigrate(&model.Attendance{})
}

func getIntArray(columnName string, array []int) string {
	var result string
	if len(array) > 0 {
		str := strings.ReplaceAll(fmt.Sprint(array), " ", ",")
		result = fmt.Sprintf(" AND %s IN (%s)", columnName, str[1:len(str)-1])
	}
	return result
}
