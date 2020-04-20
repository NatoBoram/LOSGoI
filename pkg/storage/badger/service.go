package badger

import (
	badger "github.com/dgraph-io/badger/v2"
)

// Service is a BadgerDS Service.
type Service struct {
	db *badger.DB
}

// New returns a new BadgerDS service.
func New(path string) (*Service, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}

	return &Service{
		db: db,
	}, err
}

// Close closes a DB. It's crucial to call it to ensure all the pending updates make their way to disk. Calling DB.Close() multiple times would still only close the DB once.
func (s *Service) Close() error {
	return s.db.Close()
}
