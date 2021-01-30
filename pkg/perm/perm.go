package perm

import (
	"sync"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Capacity []int
}

type Perm struct {
	check chan struct{}
	index []int
	mutex sync.Mutex

	capacity []int
}

func New(config Config) (*Perm, error) {
	if config.Capacity == nil {
		return nil, tracer.Maskf(invalidConfigError, "%T.Capacity must not be empty", config)
	}

	p := &Perm{
		check: make(chan struct{}),
		index: make([]int, len(config.Capacity)),
		mutex: sync.Mutex{},

		capacity: config.Capacity,
	}

	return p, nil
}

func (p *Perm) Check() <-chan struct{} {
	return p.check
}

func (p *Perm) Index() []int {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.index
}

func (p *Perm) Reset() {
	p.check = make(chan struct{})
	p.index = make([]int, len(p.capacity))
}

func (p *Perm) Shift() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if equal(p.capacity, p.index) {
		select {
		case <-p.Check():
			return
		default:
			close(p.check)
		}
	} else {
		p.index[len(p.index)-1]++

		for i := len(p.index) - 1; i > 0; i-- {
			if p.index[i] > p.capacity[i] {
				p.index[i] = 0
				p.index[i-1]++
			}
		}
	}
}

func equal(a []int, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
