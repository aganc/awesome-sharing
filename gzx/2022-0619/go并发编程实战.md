### RWMutex

RWMutex的方法总共有5个

![1654142321602](C:\Users\gzx\AppData\Roaming\Typora\typora-user-images\1654142321602.png)

RWMutex的实现原理

```go
type RWMutex struct {
	w Mutex // 互斥锁解决多个writer的竞争
	writerSem uint32 // writer信号量
	readerSem uint32 // reader信号量
	readerCount int32 // reader的数量
	readerWait int32 // writer等待完成的reader的数量
}
const rwmutexMaxReaders = 1 << 30
```









### WaitGroup

### Cond

### Once

### sync.Map

### Pool

### Context

### atomic

### Channel

### Semaphore

### SingleFlight

### CyclicBarrier









实践的技术分享

######  BIOS

