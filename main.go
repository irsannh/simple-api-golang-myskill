package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-api/auth"
	"simple-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type newStudent struct {
	Student_id       uint64 `json:"student_id" binding:"required"`
	Student_name     string `json:"student_name" binding:"required"`
	Student_age      uint64 `json:"student_age" binding:"required"`
	Student_address  string `json:"student_address" binding:"required"`
	Student_phone_no string `json:"student_phone_no" binding:"required"`
}

// func rowToStruct(rows *sql.Rows, dest interface{}) error {
// 	destv := reflect.ValueOf(dest).Elem()
// 	args := make([]interface{}, destv.Type().Elem().NumField())
// 	for rows.Next() {
// 		rowp := reflect.New(destv.Type().Elem())
// 		rowv := rowp.Elem()

// 		for i := 0; i < rowv.NumField(); i++ {
// 			args[i] = rowv.Field(i).Addr().Interface()
// 		}

// 		if err := rows.Scan(args...); err != nil {
// 			return err
// 		}

// 		destv.Set(reflect.Append(destv, rowv))
// 	}

// 	return nil
// }

func postHandler(ctx *gin.Context, db *gorm.DB) {
	
	// if ctx.Bind(&newStudent) == nil {
	// 	_, err := db.Exec("INSERT INTO students VALUES ($1, $2, $3, $4, $5)", newStudent.Student_id, newStudent.Student_name, newStudent.Student_age, newStudent.Student_address, newStudent.Student_phone_no)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message":"success create",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusBadRequest, gin.H{
	// 	"message": "Error",
	// })
	var newStudent newStudent

	ctx.Bind(&newStudent)
	db.Create(&newStudent)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Create",
		"data": newStudent,
	})
}

func getAllHandler(ctx *gin.Context, db *gorm.DB) {
	
	// row, err := db.Query("SELECT * FROM students")
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error":err.Error(),
	// 	})
	// }

	// rowToStruct(row, &newStudent)
	// if newStudent == nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"message": "Not Found",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": newStudent, 
	// })

	var newStudent []newStudent

	db.Find(&newStudent)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Find All",
		"data": newStudent,
	})
	
}

func getHandler(ctx *gin.Context, db *gorm.DB) {
	// var newStudent []newStudent
	// studentId := ctx.Param("student_id")
	// row, err := db.Query("SELECT * FROM students WHERE student_id = $1", studentId)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error":err.Error(),
	// 	})
	// }

	// rowToStruct(row, &newStudent)

	// if newStudent == nil {
	// 	ctx.JSON(http.StatusNotFound, gin.H{
	// 		"message": "Not Found",
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": newStudent, 
	// })

	var newStudent newStudent
	studentId := ctx.Param("student_id")
	
	if db.Find(&newStudent, "student_id=?", studentId).RecordNotFound() {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Find By Id",
		"data": newStudent,
	})
}

func putHandler(ctx *gin.Context, db *gorm.DB) {
	// studentId := ctx.Param("student_id")
	// var newStudent newStudent
	// if ctx.Bind(&newStudent) == nil {
	// 	_, err := db.Exec("UPDATE students SET student_name = $1 WHERE student_id = $2", newStudent.Student_name, studentId)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 	}

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "Success Update",
	// 	})
	// }

	
	var newStudent = newStudent{}
	studentId := ctx.Param("student_id")

	if db.Find(&newStudent, "student_id=?", studentId).RecordNotFound() {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}
	var reqStudent = newStudent
	ctx.Bind(&reqStudent)
	db.Model(&newStudent).Where("student_id=?", studentId).Update(reqStudent)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Update Data",
		"data": reqStudent,
	})
}

func deleteHandler(ctx *gin.Context, db *gorm.DB) {
	// studentId := ctx.Param("student_id")

	// _, err := db.Exec("DELETE FROM students WHERE student_id = $1", studentId)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Success Delete",
	// })

	var newStudent newStudent
	studentId := ctx.Param("student_id")
	db.Delete(&newStudent, "student_id=?", studentId)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Data",
	})
}

func setupRouter() *gin.Engine {
	conn := "postgres://postgres:postgres@127.0.0.1/postgres?sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	Migrate(db)

	r := gin.Default()

	r.POST("/login", auth.LoginHandler)

	r.POST("/student", middleware.AuthValid, func(ctx *gin.Context) {
		postHandler(ctx, db)
	})

	r.GET("/student", middleware.AuthValid, func(ctx *gin.Context) {
		getAllHandler(ctx, db)
	})

	r.GET("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		getHandler(ctx, db)
	})

	r.PUT("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		putHandler(ctx, db)
	})

	r.DELETE("/student/:student_id", middleware.AuthValid, func(ctx *gin.Context) {
		deleteHandler(ctx, db)
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run("localhost:8080")
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&newStudent{})

	data := newStudent{}
	if db.Find(&data).RecordNotFound() {
		fmt.Print("=============== run seeder ===============")
		seederUser(db)
	}
}

func seederUser(db *gorm.DB) {
	data := newStudent{
		Student_id: 1,
		Student_name: "Irsan Nur Hidayat",
		Student_age: 23,
		Student_address: "Jakarta",
		Student_phone_no: "081234567890",
	}

	db.Create(&data)
}