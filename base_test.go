package cachebenchmarks

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkBaseGCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)
	cache, err := NewWrapperGCache(loadFunc, size)
	require.NoError(b, err)

	invokeBenchmarkBase(b, keys, cache)
}

func BenchmarkBaseSCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)
	cache, err := NewWrapperSCache(loadFunc, size)
	require.NoError(b, err)
	defer cache.Close()

	invokeBenchmarkBase(b, keys, cache)
}

func BenchmarkBaseCCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)
	cache, err := NewWrapperCCache(loadFunc, size)
	require.NoError(b, err)

	invokeBenchmarkBase(b, keys, cache)
}

func BenchmarkBaseRistretto(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)

	cache, err := NewWrapperRistretto(loadFunc, size)
	require.NoError(b, err)

	invokeBenchmarkBase(b, keys, cache)
}

func invokeBenchmarkBase(b *testing.B, keys []interface{}, cache interface {
	Get(interface{}) (interface{}, error)
}) {

	gcacheVal, iGCache := cache.(*WrapperGCache)
	scacheVal, iSCache := cache.(*WrapperSCache)
	ristrettoVal, iRistretto := cache.(*WrapperRistretto)
	ccacheVal, iCcache := cache.(*WrapperCCache)

	var (
		hits  int64
		pHits = &hits
		err   error
	)

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		op := i % 3
		key := keys[i%len(keys)]

		switch op {
		case 1:
			if iGCache {
				gcacheVal.Native.Remove(key)
			} else if iSCache {
				scacheVal.Native.Del(key)
			} else if iRistretto {
				ristrettoVal.Native.Del(key)
			} else if iCcache {
				ccacheVal.Native.Delete(key.(string))
			} else {
				panic("unknown library")
			}
		default:
			_, err = cache.Get(key)
			if err == nil {
				atomic.AddInt64(pHits, 1)
			}
		}
	}

	b.StopTimer()

	var count int
	if iGCache {
		count = int(gcacheVal.Native.Len(true))
	} else if iSCache {
		count = int(scacheVal.Native.Count())
	} else if iRistretto {
		// not found method
	} else if iCcache {
		count = int(ccacheVal.Native.ItemCount())
	} else {
		panic("unknown library")
	}

	if hits < int64(b.N/2) {
		b.Error("hits", hits, "b.N", b.N)
	}

	if count < 0 {
		b.Error("count", count, b.N)
	}
}
