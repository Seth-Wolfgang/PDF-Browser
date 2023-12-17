package main


import (
	"pdfbrowser/pkg/app"
)


func main() {


	menu := app.MakeMenu()

	for true {
		menu.MainMenu()
	}

}
