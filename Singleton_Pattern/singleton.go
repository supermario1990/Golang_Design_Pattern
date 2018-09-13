// 单例实现

package singleton

import (
	"sync"
)

// Singleton is a map type
type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

// New return a singleton
func New() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})

	return instance
}
