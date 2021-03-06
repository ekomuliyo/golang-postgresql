package routes

import (
	"golang-postgresql/controllers"
	"golang-postgresql/middleware"

	"net/http"

	"github.com/labstack/echo"
)

// Init ...
func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hai, echo framework")
	})

	// image
	e.GET("/images/:param", func(c echo.Context) error {
		return c.File("images/" + c.Param("param"))
	}, middleware.AuthenticatedAsset)

	e.POST("/register_user", controllers.RegisterUser)
	e.POST("/login_user", controllers.LoginUser)

	// Group Admin
	groupAdmin := e.Group("/admin", middleware.IsAuthenticated, middleware.AdminMiddleware)
	groupAdmin.GET("/get_all_user", controllers.GetAllUser)
	groupAdmin.PUT("/update_user", controllers.UpdateUser)
	groupAdmin.DELETE("/delete_user", controllers.DeleteUser)

	groupAdmin.POST("/store_room", controllers.StoreRoom)
	groupAdmin.GET("/get_all_room", controllers.GetAllRoom)
	groupAdmin.DELETE("/delete_room", controllers.DeleteRoom)
	groupAdmin.PUT("/update_room", controllers.UpdateRoom)

	groupAdmin.GET("/get_all_booking_room", controllers.GetAllBookingRoom)
	groupAdmin.POST("/store_booking_room", controllers.StoreBookingRoom)

	// Group Guest
	groupGuest := e.Group("/guest", middleware.IsAuthenticated)
	groupGuest.GET("/get_all_room", controllers.GetAllRoom)

	return e
}
