package pidusage

import (
	"os"
	"testing"
)

var pid = os.Getpid()

func BenchmarkGetStat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStat(pid)
	}
}

// Before optimize
// $ go clean -testcache && go test -test.v -bench=BenchmarkGetStat
// goos: linux
// goarch: amd64
// pkg: github.com/struCoder/pidusage
// BenchmarkGetStat
// BenchmarkGetStat-12    	     470	   2690727 ns/op
// PASS
// ok  	github.com/struCoder/pidusage	1.533s

// After optimize
// $ go clean -testcache && go test -test.v -bench=BenchmarkGetStat
// goos: linux
// goarch: amd64
// pkg: github.com/struCoder/pidusage
// BenchmarkGetStat
// BenchmarkGetStat-12    	   28416	     36234 ns/op
// PASS
// ok  	github.com/struCoder/pidusage	1.472s
