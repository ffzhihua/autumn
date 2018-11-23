/*
@Time : 2018/11/23 14:46 
@Author : zhihua
@File : gocache
@Software: GoLand
*/
package cache

import (
	"arutam/tools/cache"
	"time"
	"sync"
)

var Cache = make(map[string]*cache.Cache)
var mutex sync.RWMutex
const (
	// For use with functions that take an expiration time.
	NoExpiration time.Duration = -1

	DefaultExpiration time.Duration = 0
)

func GetCache(table string) *cache.Cache {
	mutex.RLock()
	t, ok := Cache[table]
	mutex.RUnlock()

	if !ok {
		mutex.Lock()
		t, ok = Cache[table]
		// Double check whether the table exists or not.
		if !ok {
			t = cache.New(DefaultExpiration, 30*time.Second)
			Cache[table] = t
		}
		mutex.Unlock()
	}

	return t
}

/**
使用方式
		a := 	cache.GetCache("hello")
		a.Set("foo", "asdfasdfafasdfasdf", 5*time.Second)

		if foo, found := a.Get("foo"); found {
			fmt.Println(foo)
		}

 */