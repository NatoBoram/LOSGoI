package losgoi

// Storage is an interface used to abstract different storage implementations.
type Storage interface {
	GetDevice() (Device, error)
	SaveDevice(Device) error
	Close() error
}
