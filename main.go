package main

import (
	"flag"
	"net/http"

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
	r.GET("/shop", listObject[Shop])
	r.GET("/shop/:id", getObject[Shop])
	authorized.POST("/shop", createObject[Shop])
	authorized.POST("/shop/:id", updateObject[Shop])
	authorized.DELETE("/shop/:id", deleteObject[Shop])

	// Barber
	r.GET("/barber", listObject[Barber])
	r.GET("/barber/:id", getObject[Barber])
	authorized.POST("/barber", createObject[Barber])
	authorized.POST("/barber/:id", updateObject[Barber])
	authorized.DELETE("/barber/:id", deleteObject[Barber])

	// Service
	r.GET("/service", listObject[Service])
	r.GET("/service/:id", getObject[Service])
	authorized.POST("/service", createObject[Service])
	authorized.POST("/service/:id", updateObject[Service])
	authorized.DELETE("/service/:id", deleteObject[Service])

	// Client
	r.GET("/client", listObject[Client])
	r.GET("/client/:id", getObject[Client])
	authorized.POST("/client", createObject[Client])
	authorized.POST("/client/:id", updateObject[Client])
	authorized.DELETE("/client/:id", deleteObject[Client])

	// Attendance
	r.GET("/attendance", listObject[Attendance])
	r.GET("/attendance/:id", getObject[Attendance])
	authorized.POST("/attendance", createObject[Attendance])
	authorized.POST("/attendance/:id", updateObject[Attendance])
	authorized.DELETE("/attendance/:id", deleteObject[Attendance])
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
	db.AutoMigrate(&Shop{})
	db.AutoMigrate(&Barber{})
	db.AutoMigrate(&Service{})
	db.AutoMigrate(&Client{})
	db.AutoMigrate(&Attendance{})
}

func main() {
	dbFilename := flag.String("database", "test.db", "Path to the database file")
	flag.Parse()

	loadDatabase(*dbFilename)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
