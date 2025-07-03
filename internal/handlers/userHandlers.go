package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/TheEastWantsThis/OldNew/internal/web/users"
	userservice "github.com/TheEastWantsThis/OldNew/userService"
	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	service userservice.MainUserService
}

func NewUserHandler(s userservice.MainUserService) *UserHandlers {
	return &UserHandlers{service: s}
}

func (u *UserHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var response users.GetUsers200JSONResponse

	for _, usr := range allUsers {
		Id := int(usr.ID)
		Email := usr.Email
		Password := usr.Password
		createdAt := usr.CreatedAt
		updatedAt := usr.UpdatedAt
		var deletedAt *time.Time
		if usr.DeletedAt.Valid {
			deletedAt = &usr.DeletedAt.Time
		}

		response = append(response, users.User{
			Id:        &Id,
			Email:     &Email,
			Password:  &Password,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
			DeletedAt: deletedAt,
		})
	}
	return response, nil
}

func (u *UserHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRecuest := request.Body

	userToCreate := userservice.UsersOrm{
		Email:    *userRecuest.Email,
		Password: *userRecuest.Password,
	}
	createdUser, err := u.service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	id := int(createdUser.ID)
	createdAt := createdUser.CreatedAt
	updatedAt := createdUser.UpdatedAt
	var deletedAt *time.Time
	if createdUser.DeletedAt.Valid {
		deletedAt = &createdUser.DeletedAt.Time
	}

	response := users.PostUsers201JSONResponse{
		Id:        &id,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: deletedAt,
	}
	return response, nil
}

func (u *UserHandlers) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	update := request.Body
	id := request.Id

	updatedUser, err := u.service.UpdateUser(int(id), userservice.UsersOrm{
		Email:    *update.Email,
		Password: *update.Password,
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Could not update user")
	}

	idInt := int(id)
	createdAt := updatedUser.CreatedAt
	updatedAt := updatedUser.UpdatedAt
	var deletedAt *time.Time
	if updatedUser.DeletedAt.Valid {
		deletedAt = &updatedUser.DeletedAt.Time
	}

	result := users.User{
		Id:        &idInt,
		Email:     &updatedUser.Email,
		Password:  &updatedUser.Password,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: deletedAt,
	}

	return users.PatchUsersId200JSONResponse(result), nil

}

func (u *UserHandlers) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	if err := u.service.DeleteUser(id); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Could not delete user")
	}
	return users.DeleteUsersId204Response{}, nil
}
