package main

import "database/sql"

func selectLatest() (rows *sql.Rows, err error) {
	return db.Query("SELECT `device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs` FROM `builds_latest`;")
}

func selectOld() (rows *sql.Rows, err error) {
	return db.Query("SELECT `device`, `date`, `datetime`, `filename`, `filepath`, `sha1`, `sha256`, `size`, `type`, `version`, `ipfs` FROM `builds` WHERE `ipfs` NOT IN( SELECT `ipfs` FROM `builds_latest` );")
}
