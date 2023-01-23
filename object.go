package main

import (
	"net/http"

	"github.com/dsogari/barber-shop/graph/model"
	"github.com/gin-gonic/gin"
)

type Object interface {
	model.Shop | model.Barber | model.Service | model.Client | model.Attendance
}

func listObject[T Object](c *gin.Context) {
	var objects []T
	var err error
	if _, ok := interface{}(objects).([]model.Attendance); ok {
		err = db.Preload("Services").Find(&objects).Error
	} else {
		err = db.Find(&objects).Error
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
	if _, ok := interface{}(object).(model.Attendance); ok {
		err = db.Preload("Services").First(&object, id).Error
	} else {
		err = db.First(&object, id).Error
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
	} else if err = db.Create(&object).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}

func updateObject[T Object](c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var object1, object2 T

	id := c.Params.ByName("id")

	if err := db.First(&object1, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = c.Bind(&object2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if err = db.Model(&object1).Updates(object2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object1)
	}
}

func deleteObject[T Object](c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var object T

	id := c.Params.ByName("id")

	if err := db.Delete(&object, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}
