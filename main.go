package main

import (
	"flag"
	"net/http"

	"github.com/dsogari/barber-shop/graph/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/health", checkHealth)

	// Authorized group (uses gin.BasicAuth() middleware)
	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	// Shop
	r.GET("/shop", listObject[model.Shop])
	r.GET("/shop/:id", getObject[model.Shop])
	authorized.POST("/shop", createObject[model.Shop])
	authorized.POST("/shop/:id", updateObject[model.Shop])
	authorized.DELETE("/shop/:id", deleteObject[model.Shop])

	// Barber
	r.GET("/barber", listObject[model.Barber])
	r.GET("/barber/:id", getObject[model.Barber])
	authorized.POST("/barber", createObject[model.Barber])
	authorized.POST("/barber/:id", updateObject[model.Barber])
	authorized.DELETE("/barber/:id", deleteObject[model.Barber])

	// Service
	r.GET("/service", listObject[model.Service])
	r.GET("/service/:id", getObject[model.Service])
	authorized.POST("/service", createObject[model.Service])
	authorized.POST("/service/:id", updateObject[model.Service])
	authorized.DELETE("/service/:id", deleteObject[model.Service])

	// Client
	r.GET("/client", listObject[model.Client])
	r.GET("/client/:id", getObject[model.Client])
	authorized.POST("/client", createObject[model.Client])
	authorized.POST("/client/:id", updateObject[model.Client])
	authorized.DELETE("/client/:id", deleteObject[model.Client])

	// Attendance
	r.GET("/attendance", listObject[model.Attendance])
	r.GET("/attendance/:id", getObject[model.Attendance])
	r.POST("/query_attendance", queryAttendance)
	authorized.POST("/attendance", createObject[model.Attendance])
	authorized.POST("/attendance/:id", updateObject[model.Attendance])
	authorized.DELETE("/attendance/:id", deleteObject[model.Attendance])
	authorized.PATCH("/attendance/:id", addAttendanceService)

	return r
}

func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func loadDatabase(filename string) {
	println("Opening " + filename)
	var err error
	db, err = gorm.Open(sqlite.Open(filename+"?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Shop{})
	db.AutoMigrate(&model.Barber{})
	db.AutoMigrate(&model.Service{})
	db.AutoMigrate(&model.Client{})
	db.AutoMigrate(&model.Attendance{})
}

func main() {
	dbFilename := flag.String("database", "test.db", "Path to the database file")
	flag.Parse()

	loadDatabase(*dbFilename)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
