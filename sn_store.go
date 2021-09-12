package khl

import (
	"encoding/binary"
	"sync"

	"github.com/bits-and-blooms/bloom/v3"
)

// SnStore is the interface for storing sequence numbers.
type SnStore interface {
	TestAndInsert(int64) bool
	Lock()
	Unlock()
	Clear()
}

type bloomSnStore struct {
	filter *bloom.BloomFilter
	lock   *sync.Mutex
}

func newBloomSnStore() bloomSnStore {
	return bloomSnStore{
		filter: bloom.NewWithEstimates(1000000, 0.01),
		lock:   &sync.Mutex{},
	}
}

// TestAndInsert insert a number and check if it exists.
func (b bloomSnStore) TestAndInsert(i int64) bool {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return b.filter.TestAndAdd(buf)
}

// Lock locks the store.
func (b bloomSnStore) Lock() {
	b.lock.Lock()
}

// Unlock unlocks the store.
func (b bloomSnStore) Unlock() {
	b.lock.Unlock()
}

func (b bloomSnStore) Clear() {
	b.filter.ClearAll()
}
