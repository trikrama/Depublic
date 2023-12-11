package builder

import (
	"github.com/midtrans/midtrans-go/snap"
	repoBlog "github.com/trikrama/Depublic/internal/app/blog/repository"
	serviceBlog "github.com/trikrama/Depublic/internal/app/blog/service"
	repoEvent "github.com/trikrama/Depublic/internal/app/event/repository"
	serviceEvent "github.com/trikrama/Depublic/internal/app/event/service"
	repoNotif "github.com/trikrama/Depublic/internal/app/notification/repository"
	serviceNotif "github.com/trikrama/Depublic/internal/app/notification/service"
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
	paymentService := serviceTransaction.NewPaymentService(midtransClient)
	transactionService := serviceTransaction.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(cfg, transactionService, notifService, eventService, paymentService)

	//Blog
	blogRepository := repoBlog.NewBlogRepository(db)
	blogService := serviceBlog.NewBlogService(blogRepository)
	blogHandler := handler.NewBlogHandler(cfg, blogService)
	return router.PrivateRoutes(userHandler, eventHandler, transactionHandler, blogHandler, notifHandler)
}

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	//repository
	userRepository := repository.NewRepositoryUser(db)
	notifRepository := repoNotif.NewNotificationRepository(db)
	eventRepository := repoEvent.NewEventRepository(db)
	transactionRepository := repoTransaction.NewTransactionRepository(db)

	//service
	userService := service.NewUserService(userRepository)
	notifService := serviceNotif.NewNotificationService(notifRepository)
	eventService := serviceEvent.NewEventService(eventRepository)
	paymentService := serviceTransaction.NewPaymentService(midtransClient)
	transactionService := serviceTransaction.NewTransactionService(transactionRepository)

	//handler
	authHandler := handler.NewAuthHandler(cfg, userService, notifService)
	transactionHandler := handler.NewTransactionHandler(cfg, transactionService, notifService, eventService, paymentService)
	eventHandler := handler.NewEventHandler(cfg, eventService)
	return router.PublicRoutes(authHandler, transactionHandler, eventHandler)
}
