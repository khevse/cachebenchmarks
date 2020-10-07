package cachebenchmarks

import (
	"strconv"
	"testing"
)

type LoadFunc func(key interface{}) (value interface{}, err error)

func getPropsBenchmark(b *testing.B) (loadFunc LoadFunc, keys []interface{}, size int) {

	countOverflowKeys := 101

	size = b.N
	if size > countOverflowKeys {
		size /= 3
		size -= countOverflowKeys
	}

	keys = getKeysForBenchmark(size + countOverflowKeys)

	loadFunc = func(key interface{}) (value interface{}, err error) {
		value = key
		return
	}

	return
}

func getKeysForBenchmark(count int) (keys []interface{}) {

	keys = make([]interface{}, count)
	for i := 0; i < count; i++ {
		keys[i] = "------" + strconv.Itoa(i) // "----" - check hash function for shard id
	}

	return
}
