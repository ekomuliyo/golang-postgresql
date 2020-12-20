package routes

import (
	"golang-postgresql/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hai, echo framework")
	})

	e.POST("/login_user", controllers.LoginUser)
	e.POST("/register_user", controllers.RegisterUser)
	e.GET("/get_all_user", controllers.GetAllUser)
	e.PUT("/update_user", controllers.UpdateUser)
	e.DELETE("/delete_user", controllers.DeleteUser)

	return e
}
