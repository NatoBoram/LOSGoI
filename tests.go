package main

import (
	"fmt"
	"time"

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
			fmt.Println("Date :", time.Time(build.Date))
			fmt.Println("Datetime :", time.Time(build.Datetime))
			fmt.Println("")
		}
	}
}
