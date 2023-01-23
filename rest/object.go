package rest

import (
	"net/http"

	"github.com/dsogari/barber-shop/graph/generated"
	"github.com/dsogari/barber-shop/orm"
	"github.com/gin-gonic/gin"
)

type Object interface {
	generated.Shop | generated.Barber | generated.Service | generated.Client | generated.Attendance
}

func listObject[T Object](c *gin.Context) {
	var objects []T
	var err error
	if _, ok := interface{}(objects).([]generated.Attendance); ok {
		err = orm.Db.Preload("Services").Find(&objects).Error
	} else {
		err = orm.Db.Find(&objects).Error
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, objects)
	}
}

func getObject[T Object](c *gin.Context) {
	id := c.Params.ByName("id")

	var object T
	var err error
	if _, ok := interface{}(object).(generated.Attendance); ok {
		err = orm.Db.Preload("Services").First(&object, id).Error
	} else {
		err = orm.Db.First(&object, id).Error
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}

func createObject[T Object](c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var object T

	if err := c.Bind(&object); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if err = orm.Db.Create(&object).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}

func updateObject[T Object](c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var object1, object2 T

	id := c.Params.ByName("id")

	if err := orm.Db.First(&object1, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = c.Bind(&object2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if err = orm.Db.Model(&object1).Updates(object2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object1)
	}
}

func deleteObject[T Object](c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var object T

	id := c.Params.ByName("id")

	if err := orm.Db.Delete(&object, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}
