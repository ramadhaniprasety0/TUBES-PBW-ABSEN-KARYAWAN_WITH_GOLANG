package main

import (
	"dashboard/controllers/Absencontroller"
	"dashboard/controllers/Karywancontroller"
	"dashboard/controllers/Usercontroller"
	"dashboard/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	// Initialize the router
	r := gin.Default()

	// Connect to the database
	models.ConnectDatabase()
	db := models.DB

	// Load templates from the templates folder
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Serve static files
	r.Static("/static", "./static")

	// Initialize controllers with the database connection and templates
	Usercontroller.InitUserController(db, templates)
	Karywancontroller.InitKaryawanController(db)
	Absencontroller.InitAbsenController(db)

	// Setup session middleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// User routes
	r.GET("/register", func(c *gin.Context) {
		if err := templates.ExecuteTemplate(c.Writer, "register.html", nil); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})
	r.POST("/register", Usercontroller.RegisterUser)
	r.GET("/", func(c *gin.Context) {
		if err := templates.ExecuteTemplate(c.Writer, "login.html", nil); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})
	r.POST("/login", Usercontroller.LoginUser)
	r.GET("/logout", Usercontroller.LogoutUser)

	// Protect routes with authentication middleware
	authorized := r.Group("/")
	authorized.Use(Usercontroller.AuthRequired)
	{
		authorized.GET("/dashboard", Usercontroller.ShowDashboard)
		authorized.GET("/add", func(c *gin.Context) {
			if err := templates.ExecuteTemplate(c.Writer, "add.html", nil); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		})
		authorized.GET("/admin", func(c *gin.Context) {
			if err := templates.ExecuteTemplate(c.Writer, "admin.html", nil); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		})

		authorized.GET("/absen", func(c *gin.Context) {
			if err := templates.ExecuteTemplate(c.Writer, "absen.html", nil); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		})
		//API Data Karyawan
		authorized.GET("/api/karyawan", Karywancontroller.Index)
		authorized.GET("/api/karyawan/:id", Karywancontroller.Show)
		authorized.POST("/api/tambahkaryawan", Karywancontroller.Create)
		authorized.PUT("/api/karyawan/:id", Karywancontroller.Update)
		authorized.DELETE("/api/karyawan/:id", Karywancontroller.Delete)
		authorized.GET("/api/totalkaryawan", Karywancontroller.CountKaryawan)
		//API Data Admin
		authorized.GET("/api/admin", Usercontroller.Index)
		//API Absen Karyawan
		authorized.GET("/api/absensi", Absencontroller.Tampil)
		authorized.POST("/api/hadir", Absencontroller.Kehadiran)
	}

	// Start the server
	r.Run(":8080")
}
