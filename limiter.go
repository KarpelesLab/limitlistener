package limitlistener

import "sync"

// Limiter works kind of like the reverse of a WaitGroup. It uses a straightforward locking
// process that will cause a lock to be required for Done() but considering the locking is
// only happening on limited operations it shouldn't be slow
type Limiter struct {
	cur uint64
	max uint64
	lk  sync.Mutex
	cd  *sync.Cond
}

func NewLimiter(max uint64) *Limiter {
	res := &Limiter{max: max}
	res.cd = sync.NewCond(&res.lk)

	return res
}

// Add adds 1 to the limiter, waiting for availability if needed
func (l *Limiter) Add() {
	l.lk.Lock()
	defer l.lk.Unlock()

	for {
		if l.cur >= l.max {
			// need wait
			l.cd.Wait() // will release lock while we wait
			continue
		}
		l.cur += 1
		return
	}
}

// SetMax sets the maximum number of processes
func (l *Limiter) SetMax(newMax uint64) {
	l.lk.Lock()
	defer l.lk.Unlock()

	l.max = newMax
	l.cd.Broadcast()
}

// Done frees one entry from limiter
func (l *Limiter) Done() {
	l.lk.Lock()
	defer l.lk.Unlock()

	if l.cur == 0 {
		panic("Limiter.Done() called without matching Limited.Add()")
	}
	l.cur -= 1
	l.cd.Broadcast()
}
