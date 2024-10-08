package controller

import (
	"go-homework/internal/entity"
	"go-homework/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	TeacherService usecase.TeacherService
}

func (ctrl *TeacherController) CreateTeacher(c *gin.Context) {
	var teacher entity.Teacher

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.TeacherService.CreateTeacher(&teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, teacher)

}

func (ctrl *TeacherController) GetAllTeachers(c *gin.Context) {
	
	teachers, err := ctrl.TeacherService.GetAllTeachers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, teachers)

}

func (ctrl *TeacherController) GetTeacherByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}
	teacher, err := ctrl.TeacherService.GetTeacherByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, teacher)

}

func (ctrl *TeacherController) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teahcer ID"})
		return
	}
	var teacher entity.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teacher.ID = uint(id)
	if err := ctrl.TeacherService.UpdateTeacher(teacher.ID, &teacher); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func (ctrl *TeacherController) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}
	if err := ctrl.TeacherService.DeleteTeacher(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "Teacher is deleted"})

}
