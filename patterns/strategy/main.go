package main

import "main/pkg/strategyPkg"

func main() {
	lfu := &strategyPkg.Lfu{}
	cache := strategyPkg.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &strategyPkg.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &strategyPkg.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
