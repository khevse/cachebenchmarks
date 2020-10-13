package cachebenchmarks

import (
	"github.com/bluele/gcache"
)

type WrapperGCache struct {
	Native gcache.Cache
	Loader LoadFunc
}

func NewWrapperGCache(loadFunc LoadFunc, size int) (*WrapperGCache, error) {

	cache := gcache.New(size).LRU().LoaderFunc(gcache.LoaderFunc(loadFunc)).Build()
	return &WrapperGCache{
		Native: cache,
		Loader: loadFunc,
	}, nil
}

func (w *WrapperGCache) Get(key interface{}) (val interface{}, err error) {
	val, err = w.Native.Get(key)
	return
}
