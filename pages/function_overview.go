package pages

import "github.com/rivo/tview"

const descriptions = "Weekday Finder \nFind the weekday for any day in the past, or the future \n\nMonday Counter \nFind, for a given year, which months started on a monday"

func Overview_Page(app *tview.Application, pages *tview.Pages) tview.Primitive {
	// Creating grid to place all items
	overview_grid := tview.NewGrid()
	overview_grid.SetRows(-2, -9, -1)
	overview_grid.SetColumns(0)
	overview_grid.SetBorders(true)

	// grid elements
	overview_title := tview.NewTextView().SetText("Functions Overvierw \nCCU has many different date functions. See their descriptions here!")
	overview_descriptions := tview.NewTextView().SetText(descriptions).SetScrollable(true)
	overview_return := tview.NewButton("Return to Main Page").SetSelectedFunc(func() {
		pages.SwitchToPage("main_page")
	})

	overview_grid.AddItem(overview_title, 0, 0, 1, 1, 0, 0, false)
	overview_grid.AddItem(overview_descriptions, 1, 0, 1, 1, 0, 0, false)
	overview_grid.AddItem(overview_return, 2, 0, 1, 1, 0, 0, true)

	return overview_grid
}
