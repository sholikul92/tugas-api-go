package routes

import (
	"net/http"
	"project-api-go/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Add user
func AddUser(ctx *gin.Context) {
	var user models.Users

	// memparsing body request ke user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// jika body request tidak sesuai maka akan gagal
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Invalid input",
		})
		return
	}

	// memvalidasi ketika ada data yang tidak sesuai
	validate := validator.New()
	if err := validate.Struct(&user); err != nil {
		switch {
		case user.Nim == 0:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Gagal menambahkan data. nim tidak boleh kosong",
			})
			return
		case user.Nama == "":
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Gagal menambahkan data. nama tidak boleh kosong",
			})
			return
		case user.Jurusan == "":
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Gagal menambahkan data. jurusan tidak boleh kosong",
			})
			return
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Invalid input",
			})
			return
		}
	}

	// Jika nim sudah ada di db maka akan error
	if err := models.DB.Where("nim = ?", user.Nim).First(&user).Error; err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan data. nim sudah ada",
		})
		return
	}

	// jika berhasil mengambil body request maka akan di tambahkan kedalam db
	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data berhasil ditambahkan",
	})
}

// Get all user
func GetAllUsers(ctx *gin.Context) {

	var users []models.Users

	models.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{
		"Users": users,
	})
}

// Get user by Id
func GetUserById(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")

	models.DB.First(&user, id)
	// Jika id tidak ditemukan
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Data tidak ditemukan",
		})
		return
	}

	// Jika id berhasil ditemukan
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Update user
func UpdateUser(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")

	models.DB.First(&user, id)
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui data. id tidak ditemukan",
		})
		return
	}

	// memparsing body request ke variabel user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// jika body request tidak sesuai maka akan gagal
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Invalid input",
		})
		return
	}
	// update berhasil
	models.DB.Save(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data berhasil diperbarui",
	})
}

// Delete user
func DeleteUser(ctx *gin.Context) {
	var user models.Users
	id := ctx.Param("id")

	models.DB.First(&user, id)
	// jika id tidak ditemukan, dan data gagal dihapus
	if user.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Gagal menghapus data. id tidak ditemukan",
		})
		return
	}
	// jika data berhasil dihapus
	models.DB.Delete(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data berhasil dihapus",
	})
}
