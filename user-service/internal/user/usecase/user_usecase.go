package usecase

import "CodeMart/user-service/internal/user/entity"

type repo interface {
	Create(email, password string) (int64, error)
	GetByEmailAndPassword(email, password string) (*entity.User, error)
	GetByID(id int64) (*entity.User, error)
}

type UserUsecase struct{ repo repo }

func New(r repo) *UserUsecase { return &UserUsecase{repo: r} }

func (u *UserUsecase) Register(email, password string) (*entity.User, error) {
	id, err := u.repo.Create(email, password)
	if err != nil {
		return nil, err
	}
	return u.repo.GetByID(id)
}
func (u *UserUsecase) Login(email, password string) (*entity.User, error) {
	return u.repo.GetByEmailAndPassword(email, password)
}
func (u *UserUsecase) Get(id int64) (*entity.User, error) {
	return u.repo.GetByID(id)
}
