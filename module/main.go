package main

import (
	"fmt"
	"net/http"

	bookscontrol "module/controllers"
	Model "module/models"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Model.Task{})

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("assets"))

	router.GET("/", bookscontrol.Index)
	router.GET("/create", bookscontrol.Create)
	router.POST("/create", bookscontrol.Create)

	fmt.Println("http://LocalHost:8080")
	http.ListenAndServe(":8080", router)
}
