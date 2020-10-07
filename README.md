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
BenchmarkBaseGCache    	   25460	    140446 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache    	   25394	    141047 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache    	   25345	    136299 ns/op	      53 B/op	       1 allocs/op
BenchmarkBaseSCache
BenchmarkBaseSCache    	 2804312	       443 ns/op	      30 B/op	       0 allocs/op
BenchmarkBaseSCache    	 2912818	       524 ns/op	      38 B/op	       0 allocs/op
BenchmarkBaseSCache    	 2743322	       511 ns/op	      39 B/op	       0 allocs/op
BenchmarkBaseCCache
BenchmarkBaseCCache    	 1000000	      1205 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseCCache    	 1000000	      1340 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseCCache    	 1000000	      1302 ns/op	     160 B/op	       5 allocs/op
BenchmarkBaseRistretto
BenchmarkBaseRistretto 	 1267560	       863 ns/op	     125 B/op	       0 allocs/op
BenchmarkBaseRistretto 	 1907100	       631 ns/op	     116 B/op	       0 allocs/op
BenchmarkBaseRistretto 	 1893154	       849 ns/op	     150 B/op	       0 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	35.552s
```

```bash
go test -v -bench=BenchmarkBase -benchmem -cpu 4 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkBaseGCache
BenchmarkBaseGCache-4      	   24746	    136343 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache-4      	   24009	    132775 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseGCache-4      	   24700	    134388 ns/op	     106 B/op	       2 allocs/op
BenchmarkBaseSCache
BenchmarkBaseSCache-4      	 3302374	       415 ns/op	      37 B/op	       0 allocs/op
BenchmarkBaseSCache-4      	 3531667	       392 ns/op	      36 B/op	       0 allocs/op
BenchmarkBaseSCache-4      	 3554886	       387 ns/op	      36 B/op	       0 allocs/op
BenchmarkBaseCCache
BenchmarkBaseCCache-4      	 1299313	       892 ns/op	     153 B/op	       5 allocs/op
BenchmarkBaseCCache-4      	 1302727	       904 ns/op	     153 B/op	       5 allocs/op
BenchmarkBaseCCache-4      	 1332831	       861 ns/op	     153 B/op	       5 allocs/op
BenchmarkBaseRistretto
BenchmarkBaseRistretto-4   	 1879417	       616 ns/op	     114 B/op	       0 allocs/op
BenchmarkBaseRistretto-4   	 1879657	       852 ns/op	     133 B/op	       0 allocs/op
BenchmarkBaseRistretto-4   	 1840728	       828 ns/op	     132 B/op	       0 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	34.362s
```

## BenchmarkCuncurrent

**cache size:** b.N/3 - 101

**count keys:** cache size + 101

**operations:** set, get

```bash
go test -v -bench=BenchmarkCuncurrent -benchmem -cpu 1 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkCuncurrentSCache
BenchmarkCuncurrentSCache    	 2530086	       487 ns/op	      34 B/op	       0 allocs/op
BenchmarkCuncurrentSCache    	 2477780	       503 ns/op	      35 B/op	       0 allocs/op
BenchmarkCuncurrentSCache    	 2521921	       477 ns/op	      33 B/op	       0 allocs/op
BenchmarkCuncurrentGCache
BenchmarkCuncurrentGCache    	 2094129	       552 ns/op	      63 B/op	       1 allocs/op
BenchmarkCuncurrentGCache    	 1935363	       531 ns/op	      64 B/op	       1 allocs/op
BenchmarkCuncurrentGCache    	 2103094	       579 ns/op	      67 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto
BenchmarkCuncurrentRistretto 	 2251488	       677 ns/op	     136 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto 	 2471527	       707 ns/op	     147 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto 	 2542564	       656 ns/op	     136 B/op	       0 allocs/op
BenchmarkCuncurrentCCache
BenchmarkCuncurrentCCache    	 1000000	      1079 ns/op	     215 B/op	       7 allocs/op
BenchmarkCuncurrentCCache    	 1000000	      1095 ns/op	     215 B/op	       7 allocs/op
BenchmarkCuncurrentCCache    	 1000000	      1100 ns/op	     215 B/op	       7 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	33.387s
```

```bash
go test -v -bench=BenchmarkCuncurrent -benchmem -cpu 4 -count 3
goos: linux
goarch: amd64
pkg: github.com/khevse/cachebenchmarks
BenchmarkCuncurrentSCache
BenchmarkCuncurrentSCache-4      	 5672707	       203 ns/op	       7 B/op	       0 allocs/op
BenchmarkCuncurrentSCache-4      	 5909053	       204 ns/op	       7 B/op	       0 allocs/op
BenchmarkCuncurrentSCache-4      	 5496765	       200 ns/op	       7 B/op	       0 allocs/op
BenchmarkCuncurrentGCache
BenchmarkCuncurrentGCache-4      	 2477260	       482 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentGCache-4      	 2515357	       502 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentGCache-4      	 2496384	       489 ns/op	      15 B/op	       0 allocs/op
BenchmarkCuncurrentRistretto
BenchmarkCuncurrentRistretto-4   	 7988703	       163 ns/op	      86 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto-4   	 9735900	       150 ns/op	      85 B/op	       1 allocs/op
BenchmarkCuncurrentRistretto-4   	10090216	       149 ns/op	      85 B/op	       1 allocs/op
BenchmarkCuncurrentCCache
BenchmarkCuncurrentCCache-4      	 1805960	       665 ns/op	     183 B/op	       6 allocs/op
BenchmarkCuncurrentCCache-4      	 1867462	       672 ns/op	     182 B/op	       6 allocs/op
BenchmarkCuncurrentCCache-4      	 1806885	       611 ns/op	     184 B/op	       6 allocs/op
PASS
ok  	github.com/khevse/cachebenchmarks	24.052s
```