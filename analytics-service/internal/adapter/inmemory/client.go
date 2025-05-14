package inmemory

import (
	"log"
	"slices"
	"sync"
	"time"

	"CodeMart/analytics-service/internal/model"
)

type ViewCache struct {
	views map[int][]model.View
	mu    sync.RWMutex
}

func NewViewCache() *ViewCache {
	return &ViewCache{
		views: make(map[int][]model.View),
	}
}

func (c *ViewCache) Add(view model.View) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.views[view.ProductID] = append(c.views[view.ProductID], view)
}

func (c *ViewCache) GetByProduct(productID int) []model.View {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return append([]model.View(nil), c.views[productID]...)
}

func (c *ViewCache) GetAll() map[int][]model.View {
	c.mu.RLock()
	defer c.mu.RUnlock()
	copyMap := make(map[int][]model.View, len(c.views))
	for pid, vlist := range c.views {
		copyMap[pid] = slices.Clone(vlist)
	}
	return copyMap
}

func (c *ViewCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.views = make(map[int][]model.View)
}

func StartViewCacheRefresher(cache *ViewCache) {
	go func() {
		for {
			time.Sleep(12 * time.Hour)
			cache.Clear()
			log.Println("View cache cleared after 12 hours")
		}
	}()
}
