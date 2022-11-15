package app

import (
	"database/sql"

	categoryController "github.com/Ardnh/go-ems/controller/category"
	categoryRepository "github.com/Ardnh/go-ems/repository/category"
	categoryService "github.com/Ardnh/go-ems/service/category"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB, validate *validator.Validate) *mux.Router {
	router := mux.NewRouter()

	categoryRepository := categoryRepository.NewCategoryRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validate)
	categoryController := categoryController.NewCategoryController(categoryService)

	router.HandleFunc("/category", categoryController.Create).Methods("POST")
	router.HandleFunc("/category/{id}", categoryController.Update).Methods("PUT")
	router.HandleFunc("/category/{id}", categoryController.Delete).Methods("DELETE")
	router.HandleFunc("/category/{id}", categoryController.FindById).Methods("GET")
	router.HandleFunc("/category", categoryController.FindAll).Methods("GET")

	return router
}
