package usecase

type repo interface {
	Add(userID, productID int64) error
	Remove(userID, productID int64) error
	List(userID int64) ([]int64, error)
	Clear(userID int64) error
	Count(userID int64) (int64, error)
	Has(userID, productID int64) (bool, error)
	Replace(userID int64, productIDs []int64) error
	AddMultiple(userID int64, productIDs []int64) error
	RemoveMultiple(userID int64, productIDs []int64) error
}

type CartUsecase struct{ repo repo }

func New(r repo) *CartUsecase { return &CartUsecase{repo: r} }

func (u *CartUsecase) Add(userID, productID int64) error    { return u.repo.Add(userID, productID) }
func (u *CartUsecase) Remove(userID, productID int64) error { return u.repo.Remove(userID, productID) }
func (u *CartUsecase) List(userID int64) ([]int64, error)   { return u.repo.List(userID) }
func (u *CartUsecase) Clear(userID int64) error             { return u.repo.Clear(userID) }
func (u *CartUsecase) Count(userID int64) (int64, error)    { return u.repo.Count(userID) }
func (u *CartUsecase) Has(userID, productID int64) (bool, error) {
	return u.repo.Has(userID, productID)
}
func (u *CartUsecase) Replace(userID int64, productIDs []int64) error {
	return u.repo.Replace(userID, productIDs)
}
func (u *CartUsecase) AddMultiple(userID int64, productIDs []int64) error {
	return u.repo.AddMultiple(userID, productIDs)
}
func (u *CartUsecase) RemoveMultiple(userID int64, productIDs []int64) error {
	return u.repo.RemoveMultiple(userID, productIDs)
}
