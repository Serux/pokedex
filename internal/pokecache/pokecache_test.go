package pokecache

import (
	"slices"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(time.Second * 5)
	v := []byte{33, 33, 33}
	c.Add("a", v)
	b, _ := c.Get("a")
	if slices.Compare(v, b) != 0 {
		t.Error("data from cache different")
	}
	time.Sleep(time.Second * 5)

	_, ok := c.Get("a")

	if ok {
		t.Error("cache should have been cleaned")
	}

	//Getlocations()
}
