package main

import (
	"bre-api/addOn/colortext"
	"bre-api/config"
	fiberopen "bre-api/fiber"
	"bre-api/gorm/databases"
	"fmt"
)

func main() {

	if !config.OpenServerDataBase {
		return
	}

	if err := databases.ConnectDataBase(); err != nil {
		if err.Error() == "Error Sql lite" {
			fmt.Printf("Error It Sqlite3 but U windows???\nNow Dowload terminal linux in store (U can download Ubunut or arch)\nnow to install it open U terminal linux atfer run commad update system\n\tin Ubuntu : sudo apt update\n\tin arch : sudo pacman -Syu\nnext to install golang\n\tin ubunut : sudo api insall go\n\tin arch : sudo pacman -S go or yay -S go\nnow check go version\n\tin arch and Ubunut : go check version\n")
		}
		colortext.IsErr(err)
	}

	if !config.Openflber {
		return
	}
	fiberopen.OpenServer()

}
