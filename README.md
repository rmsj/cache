# Least Recently Used Cache Exercise

This is a simple, in memory implementation of a Least Recently Used Cache.

## Assumptions and Limitations
- For the purpose of this exercise the implementation has been simplified to integer values only.
- The cache values are kept in memory 

## API

### Put
Accepts a key and a value to be stored on the cache.
If cache is about to go over capacity, the least recently used entry is removed and the new entry is added.

### Get
Returns the value of the specified given index or `-1` if not found.

### Basic Usage Example

```go
package main

import (
	"fmt"
	"github.com/rmsj/cache"
)

func main() {

	c, err := cache.NewLRUCache(10)
	if err != nil {
		panic(err)
	}
	
	c.Put(1, 20)

	// print random first name
	fmt.Println(c.Get(1)) // should output 20

	fmt.Println(c.Get(2)) // should output -1
}
```
