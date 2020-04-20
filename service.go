package losgoi

import (
	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
	"gitlab.com/NatoBoram/LOSGoI/pkg/storage"
)

// Service is an instance of LOSGoI.
type Service struct {
	config string
	cache  string

	api     *lineageos.Service
	storage *Storage
	log     *Logger
}

// New creates a new instance of LOSGoI.
func New(
	api *lineageos.Service,
	storage *storage.Storage,
) *Service {
	service := &Service{
		api:     api,
		storage: storage,
	}

	service.start()

	return service
}

// start starts the event loop and is run automatically by `New`.
func (s *Service) start() {
	for {

		// Fetch
		start := s.log.Fetching()
		apiBuilds, err := s.api.GetBuilds()
		if err != nil {
			continue
		}
		s.log.Fetched(start, apiBuilds)
		builds := losDevices(&apiBuilds)

		// Trim

	}
}
