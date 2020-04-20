package badger

import (
	badger "github.com/dgraph-io/badger/v2"
	losgoi "gitlab.com/NatoBoram/LOSGoI"
)

// GetDevice gets a device from the datastore.
func (s *Service) GetDevice(device string) (builds *losgoi.Device, err error) {
	err = s.db.View(func(txn *badger.Txn) (err error) {

		// Get Item
		item, err := txn.Get([]byte(device))
		if err != nil {
			return err
		}

		// Get Value
		return item.Value(func(b []byte) (err error) {
			err = builds.UnmarshalGob(b)
			return
		})
	})
	return
}
