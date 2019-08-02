package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func testDate() {
	devices, err := getDevices()
	if err != nil {
		fmt.Println("Couldn't get the list of devices.")
		fmt.Println(err.Error())
		return
	}

	for _, device := range devices {
		for _, build := range device {
			fmt.Println("Device :", aurora.Green(build.Filename))
			fmt.Println("Date :", build.Date.Time)
			fmt.Println("Datetime :", build.Datetime.Time)
			fmt.Println("")
		}
	}
}
