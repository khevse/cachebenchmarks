package cachebenchmarks

import "github.com/dgraph-io/ristretto"

type WrapperRistretto struct {
	Native *ristretto.Cache
	Loader LoadFunc
}

func NewWrapperRistretto(loadFunc LoadFunc, size int) (*WrapperRistretto, error) {

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: int64(size * 10),
		MaxCost:     int64(size),
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	return &WrapperRistretto{
		Native: cache,
		Loader: loadFunc,
	}, nil
}

func (w *WrapperRistretto) Get(key interface{}) (val interface{}, err error) {
	val, ok := w.Native.Get(key)
	if !ok {
		val, err = w.Loader(key)
		if err == nil {
			w.Native.Set(key, val, 1)
		}
	}
	return
}
