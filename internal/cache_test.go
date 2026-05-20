package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(2 * time.Millisecond)
	cache.Add("https://example.com/api", []byte("test response"))
	val, exists := cache.Get("https://example.com/api")
	if string(val) != "test response" || exists != true {
		t.Error("Error getting or adding data to cache")
	}
	time.Sleep(time.Millisecond * 3)
	val, exists = cache.Get("https://example.com/api")
	if exists == true {
		t.Error("Error reaping data in cache")
	}
}
