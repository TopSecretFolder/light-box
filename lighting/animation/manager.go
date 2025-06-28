package animation

import (
	"sync"
)

var GlobalManager = Manager{}

type Queue []Animation

type Manager struct {
	Queue Queue
	mu    sync.Mutex
}

func (m *Manager) Enqueue(a Animation) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Queue = append(m.Queue, a)
}

func (m *Manager) Dequeue() (Animation, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.Queue) == 0 {
		return Animation{}, false
	}

	popped := m.Queue[0]
	if len(m.Queue) > 1 {
		m.Queue = m.Queue[1:]
	} else if len(m.Queue) == 1 {
		m.Queue = []Animation{}
	}

	return popped, true
}
