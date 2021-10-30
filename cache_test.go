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

	t.Log("Given the need to test the creation of LRU cache with different capacity")
	{
		for testID, test := range tt {
			// start with test 1
			testID++
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen creating %d value of type user.", testID, test.capacity)
				{
					_, err := cache.NewLRUCache(test.capacity)

					if !test.validCapacity {
						if err != nil {
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
