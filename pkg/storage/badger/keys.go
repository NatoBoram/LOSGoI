package badger

import (
	losgoi "gitlab.com/NatoBoram/LOSGoI"
)

const (
	prefixBuild     = "build:"
	prefixIpfsBuild = "ipfsbuild:"
	prefixDevice    = "device:"
)

func keyBuild(build *losgoi.Build) string {
	return prefixBuild + build.Build.Filename
}

func keyDevice(build *losgoi.Build) string {
	return prefixDevice + build.Device
}
