package main

import (
	"fmt"

	lrucache "github.com/anjush-bhargavan/LRU-cache/LRU_Cache"
)

func main() {
	lru := lrucache.NewLRUCache(3)

	lru.Put(1, 10)
	lru.Put(2, 20)
	lru.Put(3, 30)
	lru.Put(4, 40)

	fmt.Println(lru.Get(2))

	fmt.Println(lru.Get(1))

	fmt.Println(lru.Get(3))

	lru.Put(5,50)

	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(4))
}
