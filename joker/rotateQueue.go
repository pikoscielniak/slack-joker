package joker

import "sync"

type rotateQueue struct {
	values  []string
	sync.Mutex
	maxSize int
}

func (l *rotateQueue) add(val string) {
	l.Lock()
	defer l.Unlock()

	if len(l.values) == l.maxSize {
		l.deleteFirst()
	}
	l.values = append(l.values, val)
}

func (l *rotateQueue) deleteFirst() {
	l.values = append([]string{}, l.values[1:]...)
}

func (l *rotateQueue) hasJoke(joke string) bool {
	l.Lock()
	defer l.Unlock()
	for _, j := range l.values {
		if j == joke {
			return true
		}
	}
	return false
}

func (l *rotateQueue) pop() string {
	l.Lock()
	defer l.Unlock()
	len := len(l.values)
	if len == 0 {
		return ""
	}
	val := l.values[len-1]

	l.values = append(l.values[:len-1])
	return val
}

func newRotateQueue(maxSize int) *rotateQueue {
	return &rotateQueue{
		values:  []string{},
		maxSize: maxSize,
	}
}
