package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAttendanceService(c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var attendance Attendance
	var services []Service
	var json struct {
		ServiceIDs []uint `binding:"required"`
	}

	id := c.Params.ByName("id")

	if err := db.First(&attendance, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if len(json.ServiceIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "at least one service ID must be provided"})
	} else if err = db.Find(&services, json.ServiceIDs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = db.Model(&attendance).Association("Services").Append(&services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, attendance)
	}
}
