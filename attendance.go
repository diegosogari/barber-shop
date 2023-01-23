package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAttendanceService(c *gin.Context) {
	username := c.MustGet(gin.AuthUserKey).(string)
	println("Requested by user: ", username)

	id := c.Params.ByName("id")

	var attendance Attendance
	if err := db.First(&attendance, id).Error; err == nil {
		var json struct {
			ServiceIDs []uint `binding:"required"`
		}
		if err = c.Bind(&json); err == nil {
			var services []Service
			if err = db.Find(&services, json.ServiceIDs).Error; err == nil {
				if err = db.Model(&attendance).Association("Services").Append(&services); err == nil {
					c.JSON(http.StatusOK, attendance)
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"status": err})
				}
			} else {
				c.JSON(http.StatusNotFound, gin.H{"status": err})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": err})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	}
}