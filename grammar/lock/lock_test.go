package lock

import (
	"sync"
	"testing"
	"time"
)

const cost = time.Microsecond * 10

// ** benchmark 函数用于测试读写锁的性能 **

// 参数 b 是测试对象，rw 是读写锁，read 是读操作的次数，write 是写操作的次数
func benchmark(b *testing.B, rw RW, read, write int) {
	// 循环测试次数
	for i := 0; i < b.N; i++ {
		// 创建一个等待组，用于等待所有 goroutine 完成
		var wg sync.WaitGroup
		// 循环读操作次数
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		// 循环写操作次数
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		// 等待所有 goroutine 完成
		wg.Wait()
	}
}

func BenchmarkReadMoreMX(b *testing.B)    { benchmark(b, &Lock{}, 1000, 10) }
func BenchmarkReadMoreRWMX(b *testing.B)  { benchmark(b, &RWLock{}, 1000, 10) }
func BenchmarkWriteMoreMX(b *testing.B)   { benchmark(b, &Lock{}, 10, 1000) }
func BenchmarkWriteMoreRWMX(b *testing.B) { benchmark(b, &RWLock{}, 10, 1000) }
func BenchmarkEqualMX(b *testing.B)       { benchmark(b, &Lock{}, 500, 500) }
func BenchmarkEqualRWMX(b *testing.B)     { benchmark(b, &RWLock{}, 500, 500) }
