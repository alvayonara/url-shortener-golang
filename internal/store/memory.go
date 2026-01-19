package store

import (
	"sync"
	"url-shortener-golang/models"
)

/*
*
sync.RWMutex: (Read-Write Mutex)
- multiple reads allowed
- writes exclusive --> Lock()
- RLock() --> reading (shared/multiple)
- RUnlock()
- http servers concurrent by default

Why do we use RWMutex here?
Because Go HTTP servers are concurrent, maps are not thread-safe, and we want multiple readers but exclusive writes.
*/
type MemoryStore struct {
	mu    sync.RWMutex           // lock
	links map[string]models.Link // shared resource
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		links: make(map[string]models.Link),
	}
}

/*
*
Writing logic
 1. Lock() operation first, block other goroutines (readers also must wait)
 2. When this function exits, unlock automatically
    -> prevent forgetting unlock
    -> prevent deadlock
    -> runs even if function return early

    •	Mutex = door
    •	Lock() = close door
    •	Unlock() = open door
    •	RLock() = allow multiple readers
    •	defer = “clean up on exit”
*/
func (m *MemoryStore) Save(link models.Link) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.links[link.Code] = link
	return nil
}

/*
*
Read logic
1. Use RLock() because to allow multiple readers from any clients. (read only)
2. User RUnlock() right after reading done
*/
func (m *MemoryStore) Get(code string) (models.Link, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	link, ok := m.links[code]
	return link, ok
}
