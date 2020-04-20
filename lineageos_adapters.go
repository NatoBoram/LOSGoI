package losgoi

import (
	"time"

	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

func losDevices(dtoDevices *lineageos.Builds) []*Build {
	var builds []*Build
	for device, dtoDevice := range *dtoDevices {
		builds = append(builds, losBuilds(device, dtoDevice)...)
	}
	return builds
}

func losBuilds(device string, dtoBuilds []lineageos.Build) []*Build {
	builds := make([]*Build, len(dtoBuilds))
	for _, dtoBuild := range dtoBuilds {
		builds = append(builds, losBuild(device, &dtoBuild))
	}
	return builds
}

func losBuild(device string, losBuild *lineageos.Build) *Build {
	return &Build{
		Device: device,
		Build:  losBuild,
	}
}

func (s *Service) latest(devices *lineageos.Builds) []*Build {
	var builds []*Build
	for codename, device := range *devices {
		var latest time.Time

		// Get the max date for this device
		for _, build := range device {
			if latest.Before(time.Time(build.Datetime)) {
				latest = time.Time(build.Datetime)
			}
		}

		// Put the build corresponding to the max date in the array
		for _, build := range device {
			b := losBuild(codename, &build)
			if latest.Equal(time.Time(build.Datetime)) {
				builds = append(builds, b)
			} else {
				s.log.Skipping(b)
			}
		}

	}
	return builds
}
