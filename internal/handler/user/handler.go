package user

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type userHandler struct {
	getAllUser       user.GetAllUserUsecase
	getByIdUser      user.GetByIdUserUsecase
	getWithLimitUser user.GetWithLimitUserUsecase
	insertUser       user.InsertUserUsecase
	updateUser       user.UpdateUserUsecase
	deleteUser       user.DeleteUserUsecase
}

func NewUserHandler(
	getAllUser user.GetAllUserUsecase,
	getByIdUser user.GetByIdUserUsecase,
	getWithLimitUser user.GetWithLimitUserUsecase,
	insertUser user.InsertUserUsecase,
	updateUser user.UpdateUserUsecase,
	deleteUser user.DeleteUserUsecase,
) handlers.HandlerInterface {
	return &userHandler{
		getAllUser:       getAllUser,
		getByIdUser:      getByIdUser,
		getWithLimitUser: getWithLimitUser,
		insertUser:       insertUser,
		updateUser:       updateUser,
		deleteUser:       deleteUser,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *userHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user", HandlerFunc: h.get},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user/:id", HandlerFunc: h.getById},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user/", HandlerFunc: h.getWithLimit},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/user", HandlerFunc: h.insert},
		handlers.HandlerStruct{Method: http.MethodPut, Path: "/user/:id", HandlerFunc: h.update},
		handlers.HandlerStruct{Method: http.MethodDelete, Path: "/user/:id", HandlerFunc: h.delete},
	)
	return
}
