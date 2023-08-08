package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type userHandler struct {
	getAllUser       user.GetAllUserUsecase
	getByIdUser      user.GetByIdUserUsecase
	getWithLimitUser user.GetWithLimitUserUsecase
	updateUser       user.UpdateUserUsecase
	deleteUser       user.DeleteUserUsecase

	validate *validator.Validate
}

func NewUserHandler(
	getAllUser user.GetAllUserUsecase,
	getByIdUser user.GetByIdUserUsecase,
	getWithLimitUser user.GetWithLimitUserUsecase,
	updateUser user.UpdateUserUsecase,
	deleteUser user.DeleteUserUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &userHandler{
		getAllUser:       getAllUser,
		getByIdUser:      getByIdUser,
		getWithLimitUser: getWithLimitUser,
		updateUser:       updateUser,
		deleteUser:       deleteUser,
		validate:         validate,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *userHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user", HandlerFunc: h.get},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user/profile", HandlerFunc: h.getById},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user/", HandlerFunc: h.getWithLimit},
		handlers.HandlerStruct{Method: http.MethodPut, Path: "/user/update", HandlerFunc: h.update},
		handlers.HandlerStruct{Method: http.MethodDelete, Path: "/user/delete", HandlerFunc: h.delete},
	)
	return
}
