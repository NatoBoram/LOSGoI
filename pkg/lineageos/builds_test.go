package lineageos_test

import (
	"testing"
	"time"

	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

const buildsdto = `
{
	"sailfish": [
		{
			"date": "2020-02-04",
			"datetime": 1580774400,
			"filename": "lineage-16.0-20200204-nightly-sailfish-signed.zip",
			"filepath": "/full/sailfish/20200204/lineage-16.0-20200204-nightly-sailfish-signed.zip",
			"recovery": {
				"filename": "lineage-16.0-20200204-recovery-sailfish.img",
				"filepath": "/recovery/sailfish/20200204/lineage-16.0-20200204-recovery-sailfish.img",
				"sha1": "206d27898f9e1770cd5678da77d288d82b7a1ce6",
				"sha256": "5a29e7ba63dc4d0b7049838f6450216081376991d4908fdcf5f32083cffcffcb",
				"size": 26973480
			},
			"sha1": "8b8a28966dab3b261ed7aa07d216570a799cefd4",
			"sha256": "4279cf5b2b883309475412fcf1d762e82d344ee1d339f46f98a9c18b78fe3760",
			"size": 610956719,
			"type": "nightly",
			"version": "16.0"
		}
	]
}`

func TestJSON(t *testing.T) {
	builds, err := lineageos.UnmarshalBuilds([]byte(buildsdto))
	if err != nil {
		t.Error("Couldn't unmarshal builds.", err)
	}

	_, err = builds.Marshal()
	if err != nil {
		t.Error("Couldn't marshal builds.", err)
	}

	sailfish := builds["sailfish"]
	build := sailfish[0]
	date := time.Time(build.Date)
	datetime := time.Time(build.Datetime)

	t.Logf("date: %s, datetime: %s", date.String(), datetime.String())
}
