package lineageos_test

import (
	"strconv"
	"testing"
	"time"

	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
)

const buildDateDTO = "2020-02-04"

func TestBuildDate(t *testing.T) {
	bd := lineageos.BuildDate{}
	err := bd.UnmarshalJSON([]byte(buildDateDTO))
	if err != nil {
		t.Error("Couldn't unmarshal build date.", err)
	}

	expected, _ := time.Parse("2006-01-02", buildDateDTO)
	got := time.Time(bd)

	if got.String() != expected.String() {
		t.Errorf("Expected %s, got %s.", expected.String(), got.String())
	}
}

const buildDateTimeDTO = "1580774400"

func TestBuildDateTime(t *testing.T) {
	bdt := lineageos.BuildDateTime{}
	err := bdt.UnmarshalJSON([]byte(buildDateTimeDTO))
	if err != nil {
		t.Error("Couldn't unmarshal build date time.", err)
	}

	bdtdto, _ := strconv.ParseInt(buildDateTimeDTO, 10, 64)
	expected := time.Unix(bdtdto, 0).In(time.UTC)
	got := time.Time(bdt)

	if got.String() != expected.String() {
		t.Errorf("Expected %s, got %s.", expected.String(), got.String())
	}
}
