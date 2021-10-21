package test

import (
	"database/sql"
	"golang-api/app"
	"golang-api/controller"
	"golang-api/helper"
	"golang-api/middleware"
	"golang-api/repository"
	"golang-api/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func setUpTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(172.23.0.2:3306)/golang_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setUpRouter() http.Handler {
	db := setUpTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAutMiddleware(router)
}
