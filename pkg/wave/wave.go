package wave

import (
	"sync"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Length int
}

type Wave struct {
	check chan struct{}
	index []int
	mutex sync.Mutex

	length int
}

func New(config Config) (*Wave, error) {
	if config.Length == 0 {
		return nil, tracer.Maskf(invalidConfigError, "%T.Length must not be empty", config)
	}

	w := &Wave{
		check: make(chan struct{}),
		index: make([]int, config.Length),
		mutex: sync.Mutex{},

		length: config.Length,
	}

	return w, nil
}

func (w *Wave) Check() <-chan struct{} {
	return w.check
}

func (w *Wave) Index() []int {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.index
}

func (w *Wave) Reset() {
	w.check = make(chan struct{})
	w.index = make([]int, w.length)
}

func (w *Wave) Shift() {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.index[len(w.index)-1] == 1 {
		select {
		case <-w.Check():
			return
		default:
			close(w.check)
		}
	} else {
		var p int
		{
			var s bool
			for i := range w.index {
				if w.index[i] == 1 {
					p = i + 1
					s = true
					break
				}
			}

			if !s {
				p = 0
			}
		}

		for i := range w.index {
			w.index[i] = 0
		}

		w.index[p] = 1
	}
}
