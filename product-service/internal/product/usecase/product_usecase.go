package usecase

import "CodeMart/product-service/internal/product/entity"

type repo interface {
	GetAll() ([]*entity.Product, error)
	GetByID(id int64) (*entity.Product, error)
	GetByTag(tag string) ([]*entity.Product, error)
}

type ProductUsecase struct{ repo repo }

func New(r repo) *ProductUsecase { return &ProductUsecase{repo: r} }

func (u *ProductUsecase) All() ([]*entity.Product, error)             { return u.repo.GetAll() }
func (u *ProductUsecase) ByID(id int64) (*entity.Product, error)      { return u.repo.GetByID(id) }
func (u *ProductUsecase) ByTag(tag string) ([]*entity.Product, error) { return u.repo.GetByTag(tag) }
