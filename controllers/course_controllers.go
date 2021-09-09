package controllers

import (
	"app/configs"
	"app/models/courses"
	"net/http"

	"github.com/labstack/echo"
)

func CreateCourse(c echo.Context) error {
	var courseCreate courses.CourseRegister
	c.Bind(&courseCreate)
	var courseDB courses.Course
	courseDB.Name = courseCreate.Name
	courseDB.Category = courseCreate.Category
	courseDB.Level = courseCreate.Level
	courseDB.Price = courseCreate.Price
	courseDB.IsFree = courseCreate.IsFree
	err := configs.DB.Create(&courseDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed insert course", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponseSingle{
		true, "success insert course", courseDB,
	})
}

func GetListCourses(c echo.Context) error {
	var dataCourses []courses.Course
	err := configs.DB.Find(&dataCourses).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get course", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get user database", dataCourses,
	})
}

func GetCourseById(c echo.Context) error {
	var dataCourse []courses.Course
	courseID := c.Param("courseid")
	err := configs.DB.First(&dataCourse, courseID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get course course by id", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get cours by id", dataCourse,
	})
}

func GetCourseByFree(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Find(&dataCourse, "is_free = 1").Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get free course by id", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get free course by id", dataCourse,
	})
}

func GetCourseByPriceDesc(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Order("price desc").Find(&dataCourse, "is_free = 0").Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get course by price desc", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get course by price desc", dataCourse,
	})
}

func GetCourseByPriceAsc(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Order("price asc").Find(&dataCourse, "is_free = 0").Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get course by price asc", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get course by price asc", dataCourse,
	})
}

func GetCourseByName(c echo.Context) error {
	var dataCourse []courses.Course
	word := c.Param("word")
	err := configs.DB.Find(&dataCourse, "name LIKE ? ", "%"+word+"%").Error
	//err := configs.DB.Where("name LIKE ?", word).Find(&dataCourse).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get course by name", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get course by name", dataCourse,
	})
}

func GetCoursesCategory(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Distinct("category").Find(&dataCourse).Error
	//err := configs.DB.Where("name LIKE ?", word).Find(&dataCourse).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed get categoy course", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success get category course", dataCourse,
	})
}

func UpdateCourse(c echo.Context) error {
	var courseUpdate courses.CourseUpdater
	c.Bind(&courseUpdate)
	var courseDB courses.Course
	courseDB.Name = courseUpdate.Name
	courseDB.Category = courseUpdate.Category
	courseDB.Level = courseUpdate.Level
	courseDB.Price = courseUpdate.Price
	courseDB.IsFree = courseUpdate.IsFree
	courseID := c.Param("courseid")

	err := configs.DB.Where("id = ?", courseID).Updates(&courseDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed update course by id", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success update course by id", nil,
	})
}

func DeleteCourse(c echo.Context) error {
	var dataCourse []courses.Course
	courseID := c.Param("courseid")
	err := configs.DB.Delete(&dataCourse, courseID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed delete course database by id", nil,
		})
	}

	return c.JSON(http.StatusOK, courses.CourseResponses{
		true, "success delete course database by id", nil,
	})
}

func CountCourses(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Find(&dataCourse).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed count course", nil,
		})
	}
	countCours := configs.DB.Find(&dataCourse).RowsAffected
	return c.JSON(http.StatusOK, courses.CourseResponseCount{
		true, "success count all course", countCours,
	})
}

func CountCoursesFree(c echo.Context) error {
	var dataCourse []courses.Course
	err := configs.DB.Find(&dataCourse, "is_free = 0").Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, courses.CourseResponses{
			false, "failed count course", nil,
		})
	}
	countCours := configs.DB.Find(&dataCourse, "is_free = 0").RowsAffected
	return c.JSON(http.StatusOK, courses.CourseResponseCount{
		true, "success count free count", countCours,
	})
}
