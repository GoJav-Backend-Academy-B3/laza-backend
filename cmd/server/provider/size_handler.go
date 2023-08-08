package provider

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"

	handler "github.com/phincon-backend/laza/internal/handler/size"
	repository "github.com/phincon-backend/laza/internal/repo/size"
	usecase "github.com/phincon-backend/laza/internal/usecase/size"
)

func NewSizeHandler() handlers.HandlerInterface {
	database := db.GetPostgreSQLConnection()
	gorm := database.(*db.PsqlDB).Dbs

	sizeRepo := repository.NewSizeRepo(gorm)

	addSizeUsecase := usecase.NewAddSizeUsecaseImpl(sizeRepo)
	deleteSizeUsecase := usecase.NewDeleteSizeUsecaseImpl(sizeRepo)
	getSizeByIdUsecase := usecase.NewGetSizeByIdUsecaseImpl(sizeRepo)
	getAllSizeUsecase := usecase.NewGetAllSizeUsecaseImpl(sizeRepo)
	updateSizeUsecase := usecase.NewUpdateSizeUsecaseImpl(sizeRepo)

	return handler.NewSizeHandler("/size",
		addSizeUsecase,
		deleteSizeUsecase,
		getSizeByIdUsecase,
		getAllSizeUsecase,
		updateSizeUsecase)
}
