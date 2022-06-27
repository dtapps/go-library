package golock

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	l := Lock{}
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// Lock 上锁
func (l *Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// Unlock 解锁
func (l *Lock) Unlock() {
	l.c <- struct{}{}
}
