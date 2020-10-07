package cachebenchmarks

import (
	"sync/atomic"
	"testing"

	"github.com/bluele/gcache"
	"github.com/khevse/scache"
	"github.com/stretchr/testify/require"
)

func BenchmarkCuncurrentSCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)
	cache, err := scache.New(10, int64(size)).LRU().LoaderFunc(scache.LoadFunc(loadFunc)).Build()
	require.NoError(b, err)
	defer cache.Close()

	invokeBenchmarkCuncurrent(b, keys, cache)
}

func BenchmarkCuncurrentGCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)
	cache := gcache.New(size).LRU().LoaderFunc(gcache.LoaderFunc(loadFunc)).Build()

	invokeBenchmarkCuncurrent(b, keys, cache)
}

func BenchmarkCuncurrentRistretto(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)

	cache, err := NewWrapperRistretto(loadFunc, size)
	require.NoError(b, err)

	invokeBenchmarkCuncurrent(b, keys, cache)
}

func BenchmarkCuncurrentCCache(b *testing.B) {

	loadFunc, keys, size := getPropsBenchmark(b)

	cache, err := NewWrapperCCache(loadFunc, size)
	require.NoError(b, err)

	invokeBenchmarkCuncurrent(b, keys, cache)
}

func invokeBenchmarkCuncurrent(b *testing.B, keys []interface{}, cache interface {
	Get(interface{}) (interface{}, error)
}) {

	var (
		hits  int64
		pHits = &hits
	)

	b.SetParallelism(4)
	b.ResetTimer()
	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		var (
			i   int
			err error
		)

		for pb.Next() {
			key := keys[i%len(keys)]
			_, err = cache.Get(key)
			if err == nil {
				atomic.AddInt64(pHits, 1)
			}
			i++
		}
	})
	b.StopTimer()

	if hits < int64(b.N/2) {
		b.Error("hits", hits, "b.N", b.N)
	}
}
