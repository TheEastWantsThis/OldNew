package main

import (
	"log"

	db "github.com/TheEastWantsThis/OldNew/internal/DB"
	"github.com/TheEastWantsThis/OldNew/internal/handlers"
	taskservice "github.com/TheEastWantsThis/OldNew/internal/taskService"
	"github.com/TheEastWantsThis/OldNew/internal/web/tasks"
	"github.com/TheEastWantsThis/OldNew/internal/web/users"
	userservice "github.com/TheEastWantsThis/OldNew/userService"

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
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	userRepo := userservice.NewUserRepository(database)
	userSer := userservice.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userSer)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)
	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":9090"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
