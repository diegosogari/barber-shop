package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dsogari/barber-shop/graph/model"
	"github.com/dsogari/barber-shop/orm"
	"github.com/gin-gonic/gin"
)

func addAttendanceService(c *gin.Context) {
	println("Requested by user: ", c.MustGet(gin.AuthUserKey).(string))

	var attendance model.Attendance
	var services []model.Service
	var json struct {
		ServiceIDs []int `binding:"required"`
	}

	id := c.Params.ByName("id")

	if err := orm.Db.First(&attendance, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if len(json.ServiceIDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "at least one service ID must be provided"})
	} else if err = orm.Db.Find(&services, json.ServiceIDs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else if err = orm.Db.Model(&attendance).Association("Services").Append(&services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, attendance)
	}
}

func queryAttendance(c *gin.Context) {
	var attendances []model.Attendance
	var json struct {
		ShopIDs   []uint
		BarberIDs []uint
		ClientIDs []uint
		Begin     int
		End       int
	}

	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else if err = orm.Db.Preload("Services").Find(&attendances, "attended_at BETWEEN ? AND ?"+
		getIntArray("shop_id", json.ShopIDs)+
		getIntArray("barber_id", json.BarberIDs)+
		getIntArray("client_id", json.ClientIDs), json.Begin, json.End).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	} else {
		c.JSON(http.StatusOK, attendances)
	}
}

func getIntArray(columnName string, array []uint) string {
	var result string
	if len(array) > 0 {
		str := strings.ReplaceAll(fmt.Sprint(array), " ", ",")
		result = fmt.Sprintf(" AND %s IN (%s)", columnName, str[1:len(str)-1])
	}
	return result
}
