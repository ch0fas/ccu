package main

import (
	// "fmt"

	"ccu_chofas/pages"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main(){
	// Creating app and titles and buttons in main menu page
	app := tview.NewApplication()
	main_title_view := tview.NewTextView().SetText("Welcome to my app! \n Feel free to play around with the functions below").SetTextAlign(tview.AlignCenter)
	external_pages := tview.NewPages()

	grid := tview.NewGrid().
		SetRows(0,0,0,0,0).
		SetColumns(-2,-9,-1).SetBorders(true)
	
	// Buttons to access the different parts of the TUI
	creds_button := tview.NewButton("Credits").SetSelectedFunc(func() {
		external_pages.SwitchToPage("credits")
	})
	weekday_button := tview.NewButton("Weekday Finder").SetSelectedFunc(func() {
		external_pages.SwitchToPage("weekday_finder")
	})

	grid.AddItem(main_title_view, 0, 0, 1, 3, 0, 0, false)
	grid.AddItem(creds_button, 4,0, 1, 1, 0, 0, false)
	grid.AddItem(weekday_button, 1, 0, 1, 1, 0, 0, true)

	// Creating the external pages that can be accessed with the GUI
	external_pages.AddPage("main_page", grid, true, true)
	external_pages.AddPage("credits", pages.Credits_Form(app, external_pages), true, false)
	external_pages.AddPage("weekday_finder", pages.Weekday_Finder(app, external_pages), true, false)

	// Movement options in the grid
	clickable_elements := []tview.Primitive{weekday_button, creds_button}
	clickables_index := 0

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			clickables_index = (clickables_index + 1) % len(clickable_elements)
			app.SetFocus(clickable_elements[clickables_index])
			return nil
		}
		return event
	})

	
	if err := app.SetRoot(external_pages, true).Run(); err != nil{
		panic(err)
	}
}