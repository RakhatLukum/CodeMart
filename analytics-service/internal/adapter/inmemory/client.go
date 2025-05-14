package inmemory

import (
	"CodeMart/analytics-service/internal/model"
	"sync"
)

type Client struct {
	views map[int]model.View
	mu    sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		views: make(map[int]model.View),
	}
}

func (vc *Client) Set(view model.View) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	vc.views[view.ProductID] = view
}

func (vc *Client) SetMany(views []model.View) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	for _, view := range views {
		vc.views[view.ProductID] = view
	}
}

func (vc *Client) Get(productID int) (model.View, bool) {
	vc.mu.RLock()
	defer vc.mu.RUnlock()

	view, ok := vc.views[productID]
	return view, ok
}

func (vc *Client) Delete(productID int) {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	delete(vc.views, productID)
}
