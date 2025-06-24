package main

import (
	"log"

	db "github.com/TheEastWantsThis/OldNew/internal/DB"
	"github.com/TheEastWantsThis/OldNew/internal/handlers"
	taskservice "github.com/TheEastWantsThis/OldNew/internal/taskService"
	"github.com/TheEastWantsThis/OldNew/internal/web/tasks"

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
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":9090"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
