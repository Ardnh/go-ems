package app

import (
	"database/sql"

	categoryController "github.com/Ardnh/go-ems/controller/category"
	monitoringController "github.com/Ardnh/go-ems/controller/monitoring"
	superUserController "github.com/Ardnh/go-ems/controller/super_user"
	userController "github.com/Ardnh/go-ems/controller/user"
	categoryRepository "github.com/Ardnh/go-ems/repository/category"
	monitoringRepository "github.com/Ardnh/go-ems/repository/monitoring"
	superUserRepository "github.com/Ardnh/go-ems/repository/super_user"
	userRepository "github.com/Ardnh/go-ems/repository/user"
	categoryService "github.com/Ardnh/go-ems/service/category"
	monitoringService "github.com/Ardnh/go-ems/service/monitoring"
	superUserService "github.com/Ardnh/go-ems/service/super_user"
	userService "github.com/Ardnh/go-ems/service/user"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *mux.Router {
	router := mux.NewRouter()
	s := router.PathPrefix("/api").Subrouter()

	// Category
	categoryRepository := categoryRepository.NewCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validate)
	categoryController := categoryController.NewCategoryController(categoryService)

	// User
	userRepository := userRepository.NewUserRepository()
	userService := userService.NewUserService(userRepository, db, validate)
	userController := userController.NewUserController(userService)

	// Super User
	superUserRepository := superUserRepository.NewSuperUserRepository()
	superUserService := superUserService.NewSuperUserService(superUserRepository, db, validate)
	superUserController := superUserController.NewSuperUserController(superUserService)

	// Monitoring
	monitoringRepository := monitoringRepository.NewMonitoringRepository()
	monitoringService := monitoringService.NewMonitoringService(monitoringRepository, db, validate)
	monitoringController := monitoringController.NewMonitoringController(monitoringService)

	// Category Route
	s.HandleFunc("/category", categoryController.Create).Methods("POST")
	s.HandleFunc("/category/{id}", categoryController.Update).Methods("PUT")
	s.HandleFunc("/category/{id}", categoryController.Delete).Methods("DELETE")
	s.HandleFunc("/category/{id}", categoryController.FindById).Methods("GET")
	s.HandleFunc("/category", categoryController.FindAll).Methods("GET")

	// User Route
	s.HandleFunc("/user/register", userController.Register).Methods("POST")
	s.HandleFunc("/user", userController.Update).Methods("PUT")
	s.HandleFunc("/user/login", userController.Login).Methods("POST")

	// Super User Route
	s.HandleFunc("/super-user/register", superUserController.Register).Methods("POST")
	s.HandleFunc("/super-user/login", superUserController.Register).Methods("POST")

	// Monitoring Route
	s.HandleFunc("/monitoring/user", monitoringController.GetTotalUser).Methods("GET")
	s.HandleFunc("/monitoring/events", monitoringController.GetTotalEvents).Methods("GET")
	s.HandleFunc("/monitoring/advertise", monitoringController.GetTotalAdvertise).Methods("GET")

	return router
}
