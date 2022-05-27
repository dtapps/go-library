package gocache

// RistrettoCache https://github.com/dgraph-io/ristretto
type RistrettoCache struct {
	db              *Ristretto       // 驱动
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (r *Ristretto) NewCache() *RistrettoCache {
	return &RistrettoCache{db: r}
}

// GetInterface 缓存操作
func (rc *RistrettoCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return rc.GetterInterface()
	}

	// 如果不存在，则调用GetterInterface
	ret, found := rc.db.Get(key)

	if found == false {
		rc.db.Set(key, f(), 1)
		ret, _ = rc.db.Get(key)
	}

	return
}
