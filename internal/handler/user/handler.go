package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/middleware"
)

type userHandler struct {
	basePath           string
	getAllUser         user.GetAllUserUsecase
	getByIdUser        user.GetByIdUserUsecase
	getWithLimitUser   user.GetWithLimitUserUsecase
	updateUser         user.UpdateUserUsecase
	changePasswordUser user.ChangePasswordUserUsecase
	deleteUser         user.DeleteUserUsecase

	validate *validator.Validate
}

func NewUserHandler(
	basePath string,
	getAllUser user.GetAllUserUsecase,
	getByIdUser user.GetByIdUserUsecase,
	getWithLimitUser user.GetWithLimitUserUsecase,
	updateUser user.UpdateUserUsecase,
	changePasswordUser user.ChangePasswordUserUsecase,
	deleteUser user.DeleteUserUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &userHandler{
		basePath:           basePath,
		getAllUser:         getAllUser,
		getByIdUser:        getByIdUser,
		getWithLimitUser:   getWithLimitUser,
		updateUser:         updateUser,
		changePasswordUser: changePasswordUser,
		deleteUser:         deleteUser,
		validate:           validate,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *userHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath,
			HandlerFunc: h.get,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath + "/profile",
			HandlerFunc: h.getById,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath + "/",
			HandlerFunc: h.getWithLimit,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        h.basePath + "/update",
			HandlerFunc: h.update,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        h.basePath + "/change-password",
			HandlerFunc: h.changePassword,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        h.basePath + "/delete",
			HandlerFunc: h.delete,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
	)
	return
}
