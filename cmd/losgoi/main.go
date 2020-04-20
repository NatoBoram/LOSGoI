package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/logrusorgru/aurora"
	losgoi "gitlab.com/NatoBoram/LOSGoI"
	"gitlab.com/NatoBoram/LOSGoI/internal/self"
	"gitlab.com/NatoBoram/LOSGoI/pkg/lineageos"
	"gitlab.com/NatoBoram/LOSGoI/pkg/storage/badger"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println(aurora.Bold("LOSGoI :"), "LineageOS Goes to IPFS.")
	fmt.Println("Copyright © 2018-2020 Nato Boram")
	fmt.Println("This program is free software : you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY ; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with this program. If not, see http://www.gnu.org/licenses/.")
	fmt.Println(aurora.Bold("Contact :"), aurora.Blue("https://gitlab.com/NatoBoram/LOSGoI"))
	fmt.Println("")

	storage, err := initBadger()
	if err != nil {
		return err
	}

	los := lineageos.New(self.Agent(), &http.Client{})

	core := losgoi.New(storage, los)
}

func initBadger() (*badger.Service, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	path := configDir + os.PathSeparator + self.User + "badger"
	err = os.MkdirAll(path, losgoi.PermPrivateDirectory)
	if err != nil {
		return err
	}

	return badger.New(path)
}
