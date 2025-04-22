package usecase

type repo interface{
    Add(userID,productID int64) error
    Remove(userID,productID int64) error
    List(userID int64)([]int64,error)
}

type CartUsecase struct{ repo repo }

func New(r repo)*CartUsecase{ return &CartUsecase{repo:r}}

func (u *CartUsecase) Add(userID,productID int64) error{ return u.repo.Add(userID,productID)}
func (u *CartUsecase) Remove(userID,productID int64) error{ return u.repo.Remove(userID,productID)}
func (u *CartUsecase) List(userID int64)([]int64,error){ return u.repo.List(userID)}
