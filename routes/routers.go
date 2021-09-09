package routes

import (
	"app/constants"
	"app/controllers"
	"app/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.LogMiddleware(e)
	eAdminCourse := e.Group("/admin/course")
	eAdminCourse.GET("/lists", controllers.GetListCourses)
	eAdminCourse.POST("/create", controllers.CreateCourse)
	eAdminCourse.GET("/count", controllers.CountCourses)
	eAdminCourse.GET("/countfree", controllers.CountCoursesFree)
	eAdminCourse.DELETE("/delete/:courseid", controllers.DeleteCourse)
	eAdminCourse.PUT("/update/:courseid", controllers.UpdateCourse)
	eAdminCourse.GET("/view/:courseid", controllers.GetCourseById)

	eAdminUser := e.Group("/admin/user")
	eAdminUser.GET("/list", controllers.GetListUsers)
	eAdminUser.DELETE("/delete/:userid", controllers.DeleteUser)
	eAdminUser.GET("/count", controllers.CountUsers)

	eReg := e.Group("/register")
	eReg.POST("/register", controllers.CreateUser)

	eUser := e.Group("/user")
	eUser.GET("/login/:email/:password", controllers.LoginUser)
	eUser.Use(middleware.JWT([]byte(constants.KEY_JWT)))
	eUser.GET("/courses", controllers.GetListCourses)
	eUser.GET("/coursecategory", controllers.GetCoursesCategory)
	eUser.GET("/courses/:courseid", controllers.GetCourseById)
	eUser.GET("/free", controllers.GetCourseByFree)
	eUser.GET("/pricedesc", controllers.GetCourseByPriceDesc)
	eUser.GET("/priceasc", controllers.GetCourseByPriceAsc)
	eUser.GET("/search/:word", controllers.GetCourseByName)
	return e
}
