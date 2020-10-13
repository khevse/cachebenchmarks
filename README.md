# cachebenchmarks

## BenchmarkBase

**cache size:** b.N/3 - 101

**count keys:** cache size + 101

**operations:** set, get, del

```bash
go test -v -bench=BenchmarkBase -benchmem -cpu 1 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkBaseGCache
BenchmarkBaseGCache    	 2709675	       419 ns/op	      53 B/op	       1 allocs/op
BenchmarkBaseGCache    	 2755143	       422 ns/op	      53 B/op	       1 allocs/op
BenchmarkBaseGCache    	 2763673	       798 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseSCache
BenchmarkBaseSCache    	 2912542	       447 ns/op	      38 B/op	       0 allocs/op
BenchmarkBaseSCache    	 2867926	       464 ns/op	      38 B/op	       0 allocs/op
BenchmarkBaseSCache    	 2902165	       481 ns/op	      38 B/op	       0 allocs/op
BenchmarkBaseCCache
BenchmarkBaseCCache    	 1000000	      1173 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseCCache    	 1000000	      1243 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseCCache    	 1000000	      1175 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseRistretto
BenchmarkBaseRistretto 	 1320326	      1176 ns/op	     141 B/op	       0 allocs/op
BenchmarkBaseRistretto 	 1927710	       811 ns/op	     134 B/op	       0 allocs/op
BenchmarkBaseRistretto 	 1913476	       854 ns/op	     154 B/op	       0 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	32.180s
```

```bash
go test -v -bench=BenchmarkBase -benchmem -cpu 4 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkBaseGCache
BenchmarkBaseGCache-4      	 3305194	       567 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache-4      	 3380367	       551 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache-4      	 3324962	       345 ns/op	      53 B/op	       1 allocs/op
BenchmarkBaseSCache
BenchmarkBaseSCache-4      	 3681921	       393 ns/op	      35 B/op	       0 allocs/op
BenchmarkBaseSCache-4      	 3673528	       394 ns/op	      35 B/op	       0 allocs/op
BenchmarkBaseSCache-4      	 3656022	       388 ns/op	      35 B/op	       0 allocs/op
BenchmarkBaseCCache
BenchmarkBaseCCache-4      	 1500951	       794 ns/op	     151 B/op	       5 allocs/op
BenchmarkBaseCCache-4      	 1428800	       818 ns/op	     152 B/op	       5 allocs/op
BenchmarkBaseCCache-4      	 1465474	       788 ns/op	     151 B/op	       5 allocs/op
BenchmarkBaseRistretto
BenchmarkBaseRistretto-4   	 1883925	       609 ns/op	     116 B/op	       0 allocs/op
BenchmarkBaseRistretto-4   	 1862072	       859 ns/op	     132 B/op	       0 allocs/op
BenchmarkBaseRistretto-4   	 1879244	       862 ns/op	     133 B/op	       0 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	32.114s
```

## BenchmarkCuncurrent

**cache size:** b.N/3 - 101

**count keys:** cache size + 101

**operations:** set, get

```bash
o test -v -bench=BenchmarkCuncurrent -benchmem -cpu 1 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkCuncurrentSCache
BenchmarkCuncurrentSCache    	 2675352	       436 ns/op	      33 B/op	       0 allocs/op
BenchmarkCuncurrentSCache    	 2804212	       414 ns/op	      31 B/op	       0 allocs/op
BenchmarkCuncurrentSCache    	 2740030	       412 ns/op	      32 B/op	       0 allocs/op
BenchmarkCuncurrentGCache
BenchmarkCuncurrentGCache    	 2314141	       579 ns/op	      67 B/op	       1 allocs/op
BenchmarkCuncurrentGCache    	 2352418	       563 ns/op	      62 B/op	       1 allocs/op
BenchmarkCuncurrentGCache    	 2229130	       571 ns/op	      66 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto
BenchmarkCuncurrentRistretto 	 1850877	       660 ns/op	     123 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto 	 1911033	       677 ns/op	     121 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto 	 2622296	       776 ns/op	     148 B/op	       0 allocs/op
BenchmarkCuncurrentCCache
BenchmarkCuncurrentCCache    	 1000000	      1099 ns/op	     215 B/op	       7 allocs/op
BenchmarkCuncurrentCCache    	 1000000	      1116 ns/op	     215 B/op	       7 allocs/op
BenchmarkCuncurrentCCache    	 1000000	      1132 ns/op	     215 B/op	       7 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	32.441s
```

```bash
go test -v -bench=BenchmarkCuncurrent -benchmem -cpu 4 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkCuncurrentSCache
BenchmarkCuncurrentSCache-4      	 8165706	       150 ns/op	       9 B/op	       0 allocs/op
BenchmarkCuncurrentSCache-4      	 8594427	       146 ns/op	       9 B/op	       0 allocs/op
BenchmarkCuncurrentSCache-4      	 8506947	       149 ns/op	       9 B/op	       0 allocs/op
BenchmarkCuncurrentGCache
BenchmarkCuncurrentGCache-4      	 2521824	       451 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentGCache-4      	 2471289	       453 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentGCache-4      	 2660781	       500 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto
BenchmarkCuncurrentRistretto-4   	 9003666	       157 ns/op	      86 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto-4   	 9175309	       153 ns/op	      85 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto-4   	 9213246	       151 ns/op	      85 B/op	       1 allocs/op
BenchmarkCuncurrentCCache
BenchmarkCuncurrentCCache-4      	 2030541	       647 ns/op	     182 B/op	       6 allocs/op
BenchmarkCuncurrentCCache-4      	 1842686	       654 ns/op	     181 B/op	       6 allocs/op
BenchmarkCuncurrentCCache-4      	 1925802	       627 ns/op	     181 B/op	       6 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	26.657s
```