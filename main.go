// main.go
package main

import (
	"tugas/config"
	"tugas/handlers"
	"tugas/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// Authentication route
	r.POST("/login", handlers.Login)

	// Admin-only routes for managing data
	admin := r.Group("/admin")
	admin.Use(middleware.AdminAuth())
	admin.GET("/students", handlers.GetStudents)
	admin.POST("/students", handlers.CreateStudent)

	// Tambahkan endpoint lain untuk mentor, materi, jadwal, dan pengumuman sesuai kebutuhan
	r.Run(":8080")
}
