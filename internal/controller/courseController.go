package controller

import (
	"go-homework/internal/entity"
	"go-homework/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	CourseService usecase.CourseService
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	var course entity.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.CourseService.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)

}

func (ctrl *CourseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.CourseService.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)

}

func (ctrl *CourseController) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	course, err := ctrl.CourseService.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)

}

func (ctrl *CourseController) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	var course entity.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.ID = uint(id)
	if err := ctrl.CourseService.UpdateCourse(course.ID, &course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)

}

func (ctrl *CourseController) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	if err := ctrl.CourseService.DeleteCourse(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Course is deleted"})
}
