package storage

import (
	"context"
	"sync"
)

type MemoryStore struct {
	mutex sync.Mutex
	links map[string]*Link
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		links: make(map[string]*Link),
	}
}

func (s *MemoryStore) Save(ctx context.Context, link *Link) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.links[link.ShortLink]; ok {
		return ErrShortLinkExist
	}

	s.links[link.ShortLink] = link
	return nil
}

func (s *MemoryStore) FindByShortLink(ctx context.Context, shortLink string) (*Link, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if link, ok := s.links[shortLink]; ok {
		return link, nil
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) IncrementVisitCount(ctx context.Context, shortLink string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	link, ok := s.links[shortLink]
	if !ok {
		return ErrNotFound
	}

	link.VisitCount++
	return nil
}
