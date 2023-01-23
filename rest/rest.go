package rest

import (
	"net/http"

	"github.com/dsogari/barber-shop/graph/model"
	"github.com/gin-gonic/gin"
)

func SetupServer() (srv *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	srv = gin.Default()
	srv.GET("/health", checkHealth)

	// Authorized group (uses gin.BasicAuth() middleware)
	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	grp := srv.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	setupEndpoints(srv, grp)
	return
}

func checkHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func setupEndpoints(r *gin.Engine, g *gin.RouterGroup) {
	// Shop
	r.GET("/shop", listObject[model.Shop])
	r.GET("/shop/:id", getObject[model.Shop])
	g.POST("/shop", createObject[model.Shop])
	g.POST("/shop/:id", updateObject[model.Shop])
	g.DELETE("/shop/:id", deleteObject[model.Shop])

	// Barber
	r.GET("/barber", listObject[model.Barber])
	r.GET("/barber/:id", getObject[model.Barber])
	g.POST("/barber", createObject[model.Barber])
	g.POST("/barber/:id", updateObject[model.Barber])
	g.DELETE("/barber/:id", deleteObject[model.Barber])

	// Service
	r.GET("/service", listObject[model.Service])
	r.GET("/service/:id", getObject[model.Service])
	g.POST("/service", createObject[model.Service])
	g.POST("/service/:id", updateObject[model.Service])
	g.DELETE("/service/:id", deleteObject[model.Service])

	// Client
	r.GET("/client", listObject[model.Client])
	r.GET("/client/:id", getObject[model.Client])
	g.POST("/client", createObject[model.Client])
	g.POST("/client/:id", updateObject[model.Client])
	g.DELETE("/client/:id", deleteObject[model.Client])

	// Attendance
	r.GET("/attendance", listObject[model.Attendance])
	r.GET("/attendance/:id", getObject[model.Attendance])
	r.POST("/query_attendance", queryAttendance)
	g.POST("/attendance", createObject[model.Attendance])
	g.POST("/attendance/:id", updateObject[model.Attendance])
	g.DELETE("/attendance/:id", deleteObject[model.Attendance])
	g.PATCH("/attendance/:id", addAttendanceService)
}
