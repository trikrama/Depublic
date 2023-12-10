package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/http/handler"
)

const (
	Admin = "Admin"
	Buyer = "Buyer"
)

var (
	allRoles  = []string{Admin, Buyer}
	onlyAdmin = []string{Admin}
	onlyUser  = []string{Buyer}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PrivateRoutes(
	userHandler *handler.UserHandler, 
	eventHandler *handler.EventHandler, 
	transactionHandler *handler.TransactionHandler, 
	blogHandler *handler.BlogHandler, 
	notifHandler *handler.NotificationHandler) []*Route {
	return []*Route{
		//Router for users
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   allRoles,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   onlyAdmin,
		},

		//Router for events
		{
			Method:  echo.POST,
			Path:    "/events",
			Handler: eventHandler.CreateEvent,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/events",
			Handler: eventHandler.GetAllEvents,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/events/:id",
			Handler: eventHandler.GetEventByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.PUT,
			Path:    "/events/:id",
			Handler: eventHandler.UpdateEvent,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.DELETE,
			Path:    "/events/:id",
			Handler: eventHandler.DeleteEvent,
			Roles:   onlyAdmin,
		},

		//filter and sort
		{
			Method:  echo.GET,
			Path:    "/events/filter",
			Handler: eventHandler.GetAllEventByFilter,
			Roles:   allRoles,
		},
		//Router for transactions
		{
			Method:  echo.POST,
			Path:    "/transactions",
			Handler: transactionHandler.CreateTransaction,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/transactions",
			Handler: transactionHandler.GetTransactions,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/:id",
			Handler: transactionHandler.GetTransactionByID,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/transactions/:id",
			Handler: transactionHandler.UpdateTransaction,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.DELETE,
			Path:    "/transactions/:id",
			Handler: transactionHandler.DeleteTransaction,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/user/:id",
			Handler: transactionHandler.GetTransactionsByUser,
			Roles:   allRoles,
		},
		//Router for History Transaction
		{
		    Method:  echo.GET,
		    Path:    "/transactions/history",
		    Handler: transactionHandler.GetAllHistory,
		    Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/transactions/history/:id",
			Handler: transactionHandler.GetHistoryByUser,
			Roles:   allRoles,
		},
		//Router for blogs
		{
			Method:  echo.POST,
			Path:    "/blogs",
			Handler: blogHandler.CreateBlog,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/blogs",
			Handler: blogHandler.GetAllBlog,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/blogs/:id",
			Handler: blogHandler.GetBlogByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.PUT,
			Path:    "/blogs/:id",
			Handler: blogHandler.UpdateBlog,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.DELETE,
			Path:    "/blogs/:id",
			Handler: blogHandler.DeleteBlog,
			Roles:   onlyAdmin,
		},
		//Notification
		{
			Method:  echo.GET,
			Path:    "/notifications/user/:id",
			Handler: notifHandler.GetUserNotifications,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/notifications",
			Handler: notifHandler.GetAllNotifications,
			Roles:   onlyAdmin,
		},
	}
}

func PublicRoutes(authHandler *handler.AuthHandler, transactionHandler *handler.TransactionHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
			Roles:   allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Register,
			Roles:   onlyUser,
		},
		{
			Method:  echo.POST,
			Path:    "/webhook",
			Handler: transactionHandler.WebHookTransaction,
			Roles:   allRoles,
		},
	}
}
