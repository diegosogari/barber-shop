package resolvers

import (
	"fmt"
	"strings"

	"github.com/dsogari/barber-shop/graph/generated"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate gqlgen

type Resolver struct{}

var Db *gorm.DB

type Object interface {
	generated.Shop | generated.Service | generated.Client | generated.Barber | generated.Attendance
}

func ListObjects[T Object]() (objects []*T, err error) {
	err = Db.Preload(clause.Associations).Find(&objects).Error
	return
}

func GetObject[T Object](id int) (object *T, err error) {
	err = Db.Preload(clause.Associations).First(&object, id).Error
	return
}

func CreateObject[T Object](object *T) error {
	if err := Db.Create(object).Error; err != nil {
		return err
	}
	return Db.Preload(clause.Associations).First(object).Error
}

func UpdateObject[T Object](object *T) error {
	if err := Db.Model(object).Updates(object).Error; err != nil {
		return err
	}
	return Db.Preload(clause.Associations).First(object).Error
}

func DeleteObject[T Object](id int) (object *T, err error) {
	if err := Db.Preload(clause.Associations).First(&object, id).Error; err != nil {
		return nil, err
	}
	err = Db.Delete(&object, id).Error
	return
}

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
