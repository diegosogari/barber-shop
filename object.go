package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listObject[T Shop | Barber | Service | Client | Attendance](c *gin.Context) {
	var objects []T
	if err := db.Find(&objects).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, objects)
	}
}

func getObject[T Shop | Barber | Service | Client | Attendance](c *gin.Context) {
	id := c.Params.ByName("id")

	var object T
	if err := db.First(&object, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}

func createObject[T Shop | Barber | Service | Client | Attendance](c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	println("Requested by user: ", username)

	// Parse JSON
	var object T
	if err := c.Bind(&object); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if err = db.Create(&object).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}

func updateObject[T Shop | Barber | Service | Client | Attendance](c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	println("Requested by user: ", username)

	id := c.Params.ByName("id")

	var object1, object2 T
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

func deleteObject[T Shop | Barber | Service | Client | Attendance](c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	println("Requested by user: ", username)

	id := c.Params.ByName("id")

	var object T
	if err := db.Delete(&object, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, object)
	}
}
