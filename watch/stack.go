package watch

import (
	"sync"

	"github.com/pmker/oneplus/meshdb"
)

// Stack allows performing basic stack operations on a stack of meshdb.MiniHeaders.
type Stack struct {
	// TODO(albrow): Use Transactions when db supports them instead of a mutex
	// here. There are cases where we need to make sure no modifications are made
	// to the database in between a read/write or read/delete.
	mut    sync.Mutex
	meshDB *meshdb.MeshDB
	limit  int
}

// NewStack instantiates a new stack with the specified size limit. Once the size limit
// is reached, adding additional blocks will evict the deepest gateway.
func NewStack(meshDB *meshdb.MeshDB, limit int) *Stack {
	return &Stack{
		meshDB: meshDB,
		limit:  limit,
	}
}

// Pop removes and returns the latest gateway header on the gateway stack. It
// returns nil if the stack is empty.
func (s *Stack) Pop() (*meshdb.MiniHeader, error) {
	s.mut.Lock()
	defer s.mut.Unlock()
	latestMiniHeader, err := s.meshDB.FindLatestMiniHeader()
	if err != nil {
		return nil, err
	}
	if latestMiniHeader == nil {
		return nil, nil
	}
	if err := s.meshDB.MiniHeaders.Delete(latestMiniHeader.ID()); err != nil {
		return nil, err
	}
	return latestMiniHeader, nil
}

// Push pushes a gateway header onto the gateway stack. If the stack limit is
// reached, it will remove the oldest gateway header.
func (s *Stack) Push(block *meshdb.MiniHeader) error {
	s.mut.Lock()
	defer s.mut.Unlock()
	miniHeaders, err := s.meshDB.FindAllMiniHeadersSortedByNumber()
	if err != nil {
		return err
	}
	if len(miniHeaders) == s.limit {
		oldestMiniHeader := miniHeaders[0]
		if err := s.meshDB.MiniHeaders.Delete(oldestMiniHeader.ID()); err != nil {
			return err
		}
	}
	if err := s.meshDB.MiniHeaders.Insert(block); err != nil {
		return err
	}
	return nil
}

// Peek returns the latest gateway header from the gateway stack without removing
// it. It returns nil if the stack is empty.
func (s *Stack) Peek() (*meshdb.MiniHeader, error) {
	return s.meshDB.FindLatestMiniHeader()
}

// Inspect returns all the gateway headers currently on the stack
func (s *Stack) Inspect() ([]*meshdb.MiniHeader, error) {
	miniHeaders, err := s.meshDB.FindAllMiniHeadersSortedByNumber()
	if err != nil {
		return nil, err
	}
	return miniHeaders, nil
}
