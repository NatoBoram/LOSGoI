package badger

import (
	badger "github.com/dgraph-io/badger/v2"
	losgoi "gitlab.com/NatoBoram/LOSGoI"
)

// SaveBuild saves a build to the datastore.
func (s *Service) SaveBuild(build *losgoi.Build) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		val, err := build.MarshalGob()
		if err != nil {
			return err
		}

		txn.Set([]byte(build.Build.Filename), val)
		return nil
	})

	return err
}

// GetBuild gets a build from the datastore.
func (s *Service) GetBuild(filename string) (build *losgoi.Build, err error) {
	err = s.db.View(func(txn *badger.Txn) (err error) {

		// Get Item
		item, err := txn.Get([]byte(filename))
		if err != nil {
			return err
		}

		// Get Value
		return item.Value(func(b []byte) (err error) {
			err = build.UnmarshalGob(b)
			return
		})
	})
	return
}
