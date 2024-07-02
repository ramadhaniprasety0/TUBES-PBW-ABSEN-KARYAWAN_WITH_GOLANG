package Absencontroller

import (
	"dashboard/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func InitAbsenController(database *gorm.DB) {
	db = database
}

func Tampil(c *gin.Context) {
	var absen []models.Absensi
	db.Preload("Karyawan").Find(&absen)
	c.JSON(http.StatusOK, gin.H{"absen": absen})
}

func Kehadiran(c *gin.Context) {
	var absen models.Absensi

	if err := c.ShouldBindJSON(&absen); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Check if Karyawan exists
	var karyawan models.Karyawan
	if err := db.First(&karyawan, absen.KaryawanID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Karyawan tidak ditemukan"})
		return
	}

	// Create the attendance record
	if err := db.Create(&absen).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan absen"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Absen Hadir berhasil"})
}
