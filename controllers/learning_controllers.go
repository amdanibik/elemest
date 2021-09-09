package controllers

import (
	"app/configs"
	"app/models/learnings"
	"net/http"

	"github.com/labstack/echo"
)

func CreateLearning(c echo.Context) error {
	var lCreate learnings.LearningRegister
	c.Bind(&lCreate)
	var lDB learnings.Learning
	lDB.CourseId = lCreate.CourseId
	lDB.UserId = lCreate.UserId
	err := configs.DB.Create(&lDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, learnings.LearningResponses{
			false, "failed insert learning", nil,
		})
	}

	return c.JSON(http.StatusOK, learnings.LearningResponseSingle{
		true, "success insert user database", lDB,
	})
}

func GetListLearning(c echo.Context) error {
	var listLearning []learnings.Learning
	err := configs.DB.Find(&listLearning).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, learnings.LearningResponses{
			false, "failed get learning", nil,
		})
	}

	return c.JSON(http.StatusOK, learnings.LearningResponses{
		true, "success get learning", listLearning,
	})
}

func DeleteLearing(c echo.Context) error {
	var dataLearning []learnings.Learning
	learnID := c.Param("learnid")
	err := configs.DB.Delete(&dataLearning, learnID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, learnings.LearningResponses{
			false, "failed delete learning by id", nil,
		})
	}

	return c.JSON(http.StatusOK, learnings.LearningResponses{
		true, "success delete learning by id", nil,
	})
}
