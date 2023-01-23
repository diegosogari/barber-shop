package rest

import (
	"net/http"

	"github.com/dsogari/barber-shop/graph/generated"
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
	r.GET("/shop", listObject[generated.Shop])
	r.GET("/shop/:id", getObject[generated.Shop])
	g.POST("/shop", createObject[generated.Shop])
	g.POST("/shop/:id", updateObject[generated.Shop])
	g.DELETE("/shop/:id", deleteObject[generated.Shop])

	// Barber
	r.GET("/barber", listObject[generated.Barber])
	r.GET("/barber/:id", getObject[generated.Barber])
	g.POST("/barber", createObject[generated.Barber])
	g.POST("/barber/:id", updateObject[generated.Barber])
	g.DELETE("/barber/:id", deleteObject[generated.Barber])

	// Service
	r.GET("/service", listObject[generated.Service])
	r.GET("/service/:id", getObject[generated.Service])
	g.POST("/service", createObject[generated.Service])
	g.POST("/service/:id", updateObject[generated.Service])
	g.DELETE("/service/:id", deleteObject[generated.Service])

	// Client
	r.GET("/client", listObject[generated.Client])
	r.GET("/client/:id", getObject[generated.Client])
	g.POST("/client", createObject[generated.Client])
	g.POST("/client/:id", updateObject[generated.Client])
	g.DELETE("/client/:id", deleteObject[generated.Client])

	// Attendance
	r.GET("/attendance", listObject[generated.Attendance])
	r.GET("/attendance/:id", getObject[generated.Attendance])
	r.POST("/query_attendance", queryAttendance)
	g.POST("/attendance", createObject[generated.Attendance])
	g.POST("/attendance/:id", updateObject[generated.Attendance])
	g.DELETE("/attendance/:id", deleteObject[generated.Attendance])
	g.PATCH("/attendance/:id", addAttendanceService)
}
