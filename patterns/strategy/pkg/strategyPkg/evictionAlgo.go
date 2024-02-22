package strategyPkg

type EvictionAlgo interface {
	Evict(c *Cache)
}
