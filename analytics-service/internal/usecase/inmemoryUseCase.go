package usecase

import (
	"CodeMart/analytics-service/internal/adapter/inmemory"
	"CodeMart/analytics-service/internal/model"
)

type viewMemoryUsecase struct {
	memoryClient *inmemory.Client
}

func NewViewMemoryUsecase(memoryClient *inmemory.Client) ViewMemoryUsecase {
	return &viewMemoryUsecase{memoryClient: memoryClient}
}

func (uc *viewMemoryUsecase) Set(view model.View) {
	uc.memoryClient.Set(view)
}

func (uc *viewMemoryUsecase) SetMany(views []model.View) {
	uc.memoryClient.SetMany(views)
}

func (uc *viewMemoryUsecase) Get(productID int) (model.View, bool) {
	return uc.memoryClient.Get(productID)
}

func (uc *viewMemoryUsecase) Delete(productID int) {
	uc.memoryClient.Delete(productID)
}
