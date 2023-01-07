package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/YogeLiu/CloudDisk/pkg/util"
)

type MemoryStore struct {
	memo sync.Map
}

type item struct {
	expires int64
	value   interface{}
}

func newItem(value interface{}, expires int) item {
	expires64 := int64(expires)
	if expires > 0 {
		expires64 = time.Now().Unix() + expires64
	}
	return item{
		value:   value,
		expires: expires64,
	}
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{memo: sync.Map{}}
}

func (store *MemoryStore) Set(ctx context.Context, key string, value any, ttl int) error {
	store.memo.Store(key, newItem(value, ttl))
	return nil
}

func (store *MemoryStore) Get(ctx context.Context, key string) (any, bool) {
	return getValue(store.memo.Load(key))
}

func (store *MemoryStore) Delete(ctx context.Context, keys []string) error {
	for _, k := range keys {
		store.memo.Delete(k)
	}
	return nil
}

func getValue(data any, ok bool) (any, bool) {
	if !ok {
		return nil, ok
	}
	itemObj, ok := data.(item)
	if !ok {
		return nil, false
	}
	fmt.Printf("%d, %d", itemObj.expires, time.Now().Unix())
	if itemObj.expires > 0 && itemObj.expires < time.Now().Unix() {
		return nil, false
	}
	return itemObj.value, ok

}

func (store *MemoryStore) GarbageCollect() {
	store.memo.Range(func(key, value interface{}) bool {
		if item, ok := value.(item); ok {
			if item.expires > 0 && item.expires < time.Now().Unix() {
				util.Log().Info("Cache %q is garbage collected.", key.(string))
				store.memo.Delete(key)
			}
		}
		return true
	})
}
