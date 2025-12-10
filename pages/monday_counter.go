package pages

import (
	"strconv"
	"time"

	"github.com/rivo/tview"
)

func get_months(year_str string) string {
	year_num, _ := strconv.Atoi(year_str)
	var months []string
	for i := 1; i < 13; i++ {
		day := time.Date(year_num, time.Month(i), 1, 0, 0, 0, 0, time.UTC)
		if day.Weekday().String() == "Monday" {
			months = append(months, day.Month().String())
		}
	}

	user_answer := "For " + year_str + ", the months that start on a monday are:\n"
	for i := range len(months) {
		user_answer += months[i] + " \n"
	}

	return user_answer
}

func Monday_Counter(app *tview.Application, pages *tview.Pages) tview.Primitive {
	// Creating the grid to place all the items
	mondays_grid := tview.NewGrid()
	mondays_grid.SetRows(0, 0, 0, 0, 0)
	mondays_grid.SetColumns(0)
	mondays_grid.SetBorders(true)

	//Grid Elements
	mondays_description := tview.NewTextView().SetText("In a given year, get which months started on a monday").SetTextAlign(tview.AlignCenter)
	mondays_year_input := tview.NewInputField().SetLabel("Year").SetText("2025").SetFieldWidth(20)
	mondays_result := tview.NewTextView().SetText("Enter a year!").SetTextAlign(tview.AlignCenter)
	mondays_button := tview.NewButton("Enter").SetSelectedFunc(func() {
		year_str := mondays_year_input.GetText()
		mondays_result.SetText(get_months(year_str))
	})
	mondays_return := tview.NewButton("Return to Main Page").SetSelectedFunc(func() {
		pages.SwitchToPage("main_page")
	})

	// Place items on grid
	mondays_grid.AddItem(mondays_description, 0, 0, 1, 1, 0, 0, false)
	mondays_grid.AddItem(mondays_year_input, 1, 0, 1, 1, 0, 0, true)
	mondays_grid.AddItem(mondays_result, 2, 0, 1, 1, 0, 0, false)
	mondays_grid.AddItem(mondays_button, 3, 0, 1, 1, 0, 0, false)
	mondays_grid.AddItem(mondays_return, 4, 0, 1, 1, 0, 0, false)

	return mondays_grid

}
