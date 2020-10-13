package cachebenchmarks

import (
	"strconv"
	"strings"
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

	prefixSymbols := make([]rune, 0, 100)
	for r := 'a'; r < 'z'; r++ {
		prefixSymbols = append(prefixSymbols, r)
	}
	for r := 'A'; r < 'Z'; r++ {
		prefixSymbols = append(prefixSymbols, r)
	}
	for r := '0'; r < '9'; r++ {
		prefixSymbols = append(prefixSymbols, r)
	}

	keys = make([]interface{}, count)
	for i := 0; i < count; i++ {
		prefix := prefixSymbols[i%len(prefixSymbols)] // prefix check hash function for shard id
		keys[i] = strings.Repeat(string(prefix), 10) + strconv.Itoa(i)
	}

	return
}
