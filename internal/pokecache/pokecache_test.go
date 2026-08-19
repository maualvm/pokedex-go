package pokecache

import (
	"fmt"
	"testing"
	"time"
)

type cacheTestCase struct {
	key string
	val []byte
}

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	testCases := []cacheTestCase{
		{
			key: "https://api.example.com/v1/reource",
			val: []byte("example response"),
		},
		{
			key: "https://api.example.com",
			val: []byte("hello world!"),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(testCase.key, testCase.val)
			val, found := cache.Get(testCase.key)
			if !found {
				t.Errorf("expected to find key")
				return
			}

			if string(val) != string(testCase.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5*time.Millisecond

	cache := NewCache(interval)
	cache.Add("https://api.example.com", []byte("example response"))

	_, found := cache.Get("https://api.example.com")
	if !found {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	_, found = cache.Get("http://api.example.com")
	if found {
		t.Errorf("expected key to have been deleted")
		return
	}
}
