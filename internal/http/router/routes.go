package router

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PrivateRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		// {
		// 	Method:  echo.GET,
		// 	Path:    "/users",
		// 	Handler: userHandler.GetAllUsers,
		// },
		// {
		// 	Method:  echo.GET,
		// 	Path:    "/users/:id",
		// 	Handler: userHandler.GetUserByID,
		// },
		// {
		// 	Method:  echo.POST,
		// 	Path:    "/users",
		// 	Handler: userHandler.CreateUser,
		// },
		// {
		// 	Method:  echo.PUT,
		// 	Path:    "/users/:id",
		// 	Handler: userHandler.UpdateUser,
		// },
		// {
		// 	Method:  echo.DELETE,
		// 	Path:    "/users/:id",
		// 	Handler: userHandler.DeleteUser,
		// },
	}
}

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	// return []*Route{
	// 	{
	// 		Method:  echo.POST,
	// 		Path:    "/login",
	// 		Handler: authHandler.Login,
	// 	},
	// }
}
