package routes

import (
	"go-homework/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine, teacherController *controller.TeacherController, courseController *controller.CourseController) *gin.Engine {

	teacherRoutes := r.Group("/teachers")
	{
		teacherRoutes.POST("", teacherController.CreateTeacher)        
		teacherRoutes.GET("", teacherController.GetAllTeachers)       
		teacherRoutes.GET("/:id", teacherController.GetTeacherByID)    
		teacherRoutes.PUT("/:id", teacherController.UpdateTeacher)     
		teacherRoutes.DELETE("/:id", teacherController.DeleteTeacher)  
	}

	courseRoutes := r.Group("/courses") 
	{
		courseRoutes.POST("", courseController.CreateCourse)
		courseRoutes.GET("", courseController.GetAllCourses)
		courseRoutes.GET("/:id", courseController.GetCourseByID)
		courseRoutes.PUT("/:id", courseController.UpdateCourse)
		courseRoutes.DELETE("/:id", courseController.DeleteTeacher)
	}

	return r
}
