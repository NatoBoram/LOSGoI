package main

import (
	"database/sql"
	"fmt"
	"time"
)

func getBuildsFromRows(rows *sql.Rows) (bhs []*BuildHash, err error) {
	for rows.Next() {

		// Prevent null pointers
		bh := BuildHash{Build: &Build{}}
		var t1 time.Time
		var t2 time.Time

		// Scan
		err := rows.Scan(&bh.Build.Device, &t1, &t2, &bh.Build.Filename, &bh.Build.Filepath, &bh.Build.Sha1, &bh.Build.Sha256, &bh.Build.Size, &bh.Build.Type, &bh.Build.Version, &bh.IPFS)
		if err != nil {
			fmt.Println("Couldn't scan a row of builds.")
			fmt.Println(err.Error())
			continue
		}

		// Dates
		bh.Build.Date = BuildDate{t1}
		bh.Build.Datetime = BuildDateTime{t2}
		bhs = append(bhs, &bh)
	}

	err = rows.Err()
	return
}

func getLatestBuilds() (bhs []*BuildHash, err error) {
	rows, err := selectLatest()
	if err != nil {
		fmt.Println("Couldn't select the latest builds.")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	return getBuildsFromRows(rows)
}

func getOldBuilds() (bhs []*BuildHash, err error) {
	rows, err := selectOld()
	if err != nil {
		fmt.Println("Couldn't select older builds.")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	return getBuildsFromRows(rows)
}

func getLatestBuildsFromDevice(bh *BuildHash) (bhs []*BuildHash, err error) {
	rows, err := selectLatestFromDevice(bh)
	if err != nil {
		fmt.Println("Couldn't select the latest builds from this device.")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	return getBuildsFromRows(rows)
}

func getOldBuildsFromDevice(bh *BuildHash) (bhs []*BuildHash, err error) {
	rows, err := selectOldFromDevice(bh)
	if err != nil {
		fmt.Println("Couldn't select older builds from this device.")
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	return getBuildsFromRows(rows)
}
