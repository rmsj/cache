package cache_test

import (
	"testing"

	"github.com/rmsj/cache"
	"github.com/rmsj/cache/tests"
)

type user struct {
	firstName string
	lastName  string
	email     string
}

func TestNewLRUCache(t *testing.T) {

	tt := []struct {
		name          string
		capacity      int
		validCapacity bool
	}{
		{"Negative capacity", -1, false},
		{"Single capacity", 1, false},
		{"Valid capacity", 50, true},
	}

	t.Log(tests.Given("Given the need to test the creation of LRU cache with different capacity"))
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating %d capacity cache.", testID, test.capacity)
				{
					_, err := cache.NewLRUCache(test.capacity)

					if !test.validCapacity {
						if err == nil {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have not created cache with capacity %d: its an invalid capacity."), testID, test.capacity)
						}
						t.Logf(tests.Success("\t", "Test %d:\tShould not create cache with invalid capacity (%d): %v"), testID, test.capacity, err)
					} else {
						if err != nil {
							t.Fatalf(tests.Failed("\t", "Test %d:\tShould have created cache with capacity %d: %v."), testID, test.capacity, err)
						}
						t.Logf(tests.Success("\t", "Test %d:\tShould create cache with valid capacity (%d)"), testID, test.capacity)
					}
				}
			}

			t.Run(test.name, tf)

		}
	}
}

func TestCache(t *testing.T) {

	testID := 0
	t.Log(tests.Given("Given the need to test adding and retrieving values on the LRU Cache"))
	{
		// start with test 1
		testID++
		t.Logf("\tTest %d:\tWhen not going over the capacity", testID)
		{
			c, err := cache.NewLRUCache(4)
			if err != nil {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould have not created cache with capacity %d: its an invalid capacity."), testID)
			}

			values := []int{10, 20, 30, 40}
			// add all values to
			for i, v := range values {
				c.Put(i, v)
			}

			for i, v := range values {
				fromCache := c.Get(i)
				if v != fromCache {
					t.Fatalf(tests.Failed("\t", "Test %d:\tExpected %d got %d from cache."), testID, v, fromCache)
				}
				if idx := tests.InArray(fromCache, values); idx == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tValue %d should be in cache"), testID, v)
				}
			}

			t.Logf(tests.Success("\t", "Test %d:\tAll values added and retained in the cache"), testID)
		}

		testID++
		t.Logf("\tTest %d:\tWhen going over the capacity", testID)
		{
			c, err := cache.NewLRUCache(4)
			if err != nil {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould have not created cache with capacity %d: its an invalid capacity."), testID)
			}

			values := []int{10, 20, 30, 40}
			// add all values to
			for i, v := range values {
				c.Put(i+1, v)
			}

			for i, v := range values {
				fromCache := c.Get(i + 1)
				if v != fromCache {
					t.Fatalf(tests.Failed("\t", "Test %d:\tExpected %d got %d from cache."), testID, v, fromCache)
				}
				if idx := tests.InArray(fromCache, values); idx == -1 {
					t.Fatalf(tests.Failed("\t", "Test %d:\tValue %d should be in cache"), testID, v)
				}
			}

			// index 1 should be removed
			c.Put(5, 50)
			val := c.Get(1)
			if val != -1 {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould have removed index 1 as being least recently used."), testID)
			}

			// use index 2 - now most recently used
			val = c.Get(2)
			// should remove index 3 -> 30 as the last used
			c.Put(6, 60)
			val = c.Get(3)
			if val != -1 {
				t.Fatalf(tests.Failed("\t", "Test %d:\tShould have removed index 3 as being least recently used."), testID)
			}

			t.Logf(tests.Success("\t", "Test %d:\tAll values added and the LRU logic respected"), testID)
		}
	}
}
