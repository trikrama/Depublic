package builder

import (
	"github.com/midtrans/midtrans-go/snap"
	repoNotif "github.com/trikrama/Depublic/internal/app/notification/repository"
	serviceNotif "github.com/trikrama/Depublic/internal/app/notification/service"
	repoBlog "github.com/trikrama/Depublic/internal/app/blog/repository"
	serviceBlog "github.com/trikrama/Depublic/internal/app/blog/service"
	repoEvent "github.com/trikrama/Depublic/internal/app/event/repository"
	serviceEvent "github.com/trikrama/Depublic/internal/app/event/service"
	repoTransaction "github.com/trikrama/Depublic/internal/app/transaction/repository"
	serviceTransaction "github.com/trikrama/Depublic/internal/app/transaction/service"
	"github.com/trikrama/Depublic/internal/app/user/repository"
	"github.com/trikrama/Depublic/internal/app/user/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/handler"
	"github.com/trikrama/Depublic/internal/http/router"
	"gorm.io/gorm"
)

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	//Notification
	notifRepository := repoNotif.NewNotificationRepository(db)
	notifService := serviceNotif.NewNotificationService(notifRepository)
	notifHandler := handler.NewNotificationHandler(cfg, notifService)
	// User
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(cfg, userService, notifService)

	// Event
	eventRepository := repoEvent.NewEventRepository(db)
	eventService := serviceEvent.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(cfg, eventService)

	// Transaction
	transactionRepository := repoTransaction.NewTransactionRepository(db)
	transactionService := serviceTransaction.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(cfg, transactionService)

	//Blog
	blogRepository := repoBlog.NewBlogRepository(db)
	blogService := serviceBlog.NewBlogService(blogRepository)
	blogHandler := handler.NewBlogHandler(cfg, blogService)
	return router.PrivateRoutes(userHandler, eventHandler, transactionHandler, blogHandler, notifHandler)
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewUserService(userRepository)
	notifRepository := repoNotif.NewNotificationRepository(db)
	notifService := serviceNotif.NewNotificationService(notifRepository)
	userHandler := handler.NewUserHandler(cfg, userService, notifService)

	//Transaction
	transactionRepository := repoTransaction.NewTransactionRepository(db)
	transactionService := serviceTransaction.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(cfg, transactionService)
	return router.PublicRoutes(userHandler, transactionHandler)
}
