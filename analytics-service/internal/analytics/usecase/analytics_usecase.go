package usecase

type repo interface{
    Inc(productID int64) error
    Top(limit int)([]int64,error)
}

type AnalyticsUsecase struct{ repo repo }

func New(r repo)*AnalyticsUsecase{ return &AnalyticsUsecase{repo:r}}

func (u *AnalyticsUsecase) Log(productID int64) error{ return u.repo.Inc(productID)}
func (u *AnalyticsUsecase) Top5()([]int64,error){ return u.repo.Top(5)}
