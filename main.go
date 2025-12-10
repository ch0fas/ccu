package main

import (
	// "fmt"

	"github.com/rivo/tview"
	"ccu_chofas/pages"
)

func main(){
	// Creating app and titles and buttons in main menu page
	app := tview.NewApplication()
	main_title_view := tview.NewTextView().SetText("Welcome to my app! \n Feel free to play around with the functions below").SetTextAlign(tview.AlignCenter)
	external_pages := tview.NewPages()

	grid := tview.NewGrid().
		SetRows(0,0,0,0,0).
		SetColumns(0,0,0).SetBorders(true)
	
	// Buttons to access the different parts of the TUI
	creds_button := tview.NewButton("Credits").SetSelectedFunc(func() {
		external_pages.SwitchToPage("credits")
	})

	grid.AddItem(main_title_view, 1, 1, 1, 3, 0, 0, false)
	grid.AddItem(creds_button, 5,1, 1, 1, 0, 0, true)

	// Creating the external pages that can be accessed with the GUI
	external_pages.AddPage("main_page", grid, true, true)
	external_pages.AddPage("credits", pages.Credits_Form(app, external_pages), true, false)
	//external_pages.AddPage("weekday_finder", pages.Weekday_Finder(app, external_pages), true, false)

	
	if err := app.SetRoot(external_pages, true).Run(); err != nil{
		panic(err)
	}
}