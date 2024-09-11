package main

import (
	"go-homework/internal/controller"
	"go-homework/internal/database"
	"go-homework/internal/repository"
	"go-homework/internal/usecase"
	"go-homework/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    database.InitDB() // Initialize the DB connection first

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()

    teacherRepo := repository.NewTeacherRepo(database.DB)
    teacherService := usecase.NewTeacherService(teacherRepo)
    teacherController := &controller.TeacherController{TeacherService: *teacherService}

    courseRepo := repository.NewCourseRepo(database.DB)
    courseService := usecase.NewCourseService(courseRepo)
    courseController := &controller.CourseController{CourseService: *courseService}

    r := routes.SetUpRouter(router, teacherController, courseController)
    r.SetTrustedProxies(nil)

    r.Run(":8080")
}

