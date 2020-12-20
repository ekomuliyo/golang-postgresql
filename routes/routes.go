package routes

import (
	"golang-postgresql/controllers"
	"golang-postgresql/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hai, echo framework")
	})

	e.POST("/register_user", controllers.RegisterUser)
	e.POST("/login_user", controllers.LoginUser)
	e.GET("/get_all_user", controllers.GetAllUser, middleware.IsAuthenticated)
	e.PUT("/update_user", controllers.UpdateUser, middleware.IsAuthenticated)
	e.DELETE("/delete_user", controllers.DeleteUser, middleware.IsAuthenticated)

	return e
}
