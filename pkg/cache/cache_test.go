package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	container := struct{ Name string }{Name: "name"}
	asserts := assert.New(t)
	store := NewMemoryStore()
	testCase := map[string]interface{}{
		"t1": "str",
		"t2": 2,
		"t3": 3.5,
		"t4": time.Now(),
		"t5": []int{1, 2, 3},
		"t6": container,
	}
	for k, v := range testCase {
		asserts.NoError(store.Set(context.Background(), k, v, 20))
	}
	for k := range testCase {
		data, ok := store.Get(context.Background(), k)
		fmt.Printf("data: %v, exsit: %v\n", data, ok)
	}
	time.Sleep(time.Second * 10)
	for k := range testCase {
		data, ok := store.Get(context.Background(), k)
		fmt.Printf("data: %v, exsit: %v\n", data, ok)
	}
}
