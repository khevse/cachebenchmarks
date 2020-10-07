package cachebenchmarks

import "github.com/karlseguin/ccache/v2"

type WrapperCCache struct {
	Native *ccache.Cache
	Loader LoadFunc
}

func NewWrapperCCache(loadFunc LoadFunc, size int) (*WrapperCCache, error) {

	cache := ccache.New(ccache.Configure().MaxSize(int64(size)).ItemsToPrune(10))

	return &WrapperCCache{
		Native: cache,
		Loader: loadFunc,
	}, nil
}

func (w *WrapperCCache) Get(key interface{}) (val interface{}, err error) {
	k := key.(string)
	item := w.Native.Get(k)
	if item == nil || item.Expired() {
		val, err = w.Loader(k)
		if err == nil {
			w.Native.Set(k, val, 0)
		}
	} else {

		val = item.Value()
	}

	return
}
