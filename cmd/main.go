package main

import (
	db "dacode/OldNewProdleckt/internal/DB"
	"dacode/OldNewProdleckt/internal/handlers"
	taskservice "dacode/OldNewProdleckt/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Coud not connect to DB: %v", err)
	}

	e := echo.New()

	taskRepo := taskservice.NewTaskRepository(database)

	//idGen := &taskservice.AutoIncrement{}
	taskService := taskservice.NewTaskService(taskRepo) //, idGen
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", taskHandlers.GetHandler)
	e.POST("/task", taskHandlers.PostHandler)
	e.PATCH("/task/:ID", taskHandlers.PatchHandler)
	e.DELETE("/task/:ID", taskHandlers.DeleteHandler)
	e.Start("localhost:9090")
}
