package lock

import (
	"sync"
	"testing"
	"time"
)

const cost = time.Microsecond * 10

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkReadMoreMX(b *testing.B)    { benchmark(b, &Lock{}, 100, 1) }
func BenchmarkReadMoreRWMX(b *testing.B)  { benchmark(b, &RWLock{}, 100, 1) }
func BenchmarkWriteMoreMX(b *testing.B)   { benchmark(b, &Lock{}, 10, 100) }
func BenchmarkWriteMoreRWMX(b *testing.B) { benchmark(b, &RWLock{}, 10, 100) }
func BenchmarkEqualMX(b *testing.B)       { benchmark(b, &Lock{}, 50, 50) }
func BenchmarkEqualRWMX(b *testing.B)     { benchmark(b, &RWLock{}, 50, 50) }
