package cachebenchmarks

import (
	"github.com/khevse/scache"
)

type WrapperSCache struct {
	Native *scache.Cache
	Loader LoadFunc
}

// Use wrapper because default load function with mutex
func NewWrapperSCache(loadFunc LoadFunc, size int) (*WrapperSCache, error) {

	cache, err := scache.New(100, int64(size)).
		ItemsToPrune(10).
		LRU().
		LoaderFunc(scache.LoadFunc(loadFunc)).
		Build()
	if err != nil {
		return nil, err
	}
	return &WrapperSCache{
		Native: cache,
		Loader: loadFunc,
	}, nil
}

func (w *WrapperSCache) Get(key interface{}) (val interface{}, err error) {

	item, err := w.Native.Get(key)
	if err == scache.ErrNotFound {
		val, err = w.Loader(key)
		if err == nil {
			w.Native.Set(key, val)
		}

	} else {
		val = item
	}

	return
}

func (w *WrapperSCache) Close() {
	w.Native.Close()
}
