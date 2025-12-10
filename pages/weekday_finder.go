package pages

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func obtain_weekday(day_str string, month_str string, year_str string) string {
	day_num,_ := strconv.Atoi(day_str)
	month_num,_ := strconv.Atoi(month_str)
	year_num,_ := strconv.Atoi(year_str)

	user_date := time.Date(year_num, time.Month(month_num), day_num, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	weekday := user_date.Weekday().String()

	if user_date.Before(today){
		answer :=  day_str + "/" + month_str + "/" + year_str + " was a " + weekday
		return answer
	} else {
		answer :=  day_str + "/" + month_str + "/" + year_str + " will be a " + weekday
		return answer
	}
	
}

func Weekday_Finder(app *tview.Application, pages *tview.Pages) tview.Primitive {	
	// Creating a grid to place all the items
	weekday_layout := tview.NewGrid()
	weekday_layout.SetRows(-2, -9, -1)
	weekday_layout.SetColumns(0, 0, 0)
	weekday_layout.SetBorders(true).SetTitle("Foo")

	// grid elements
	weekday_description := tview.NewTextView().SetText("Input any date and get the weekday for that date \nDates are represented in a DD/MM/YYYY format!").SetTextAlign(tview.AlignCenter)
	weekday_form := tview.NewForm()
	weekday_return := tview.NewButton("Return to Main Page").SetSelectedFunc(func() {
		pages.SwitchToPage("main_page")
	})
	weekday_result := tview.NewTextView().SetText("")

	grid_elements := []tview.Primitive{weekday_form, weekday_return}
	current_focus := 0

	// Configuration of form
	day_field := tview.NewInputField().SetLabel("Day").SetText("1").SetFieldWidth(20)
	month_field := tview.NewDropDown().SetLabel("Month").SetOptions([]string{"01","02","03","04","05","06","07","08","09","10","11","12"}, nil).SetCurrentOption(0).SetFieldWidth(20)
	year_field := tview.NewInputField().SetLabel("Year").SetText("2025").SetFieldWidth(20)

	weekday_form.AddFormItem(day_field)
	weekday_form.AddFormItem(month_field)
	weekday_form.AddFormItem(year_field)


	weekday_form.AddButton("Calculate Result", func(){
		day_str := day_field.GetText()
		_, month_str := month_field.GetCurrentOption()
		year_str := year_field.GetText()
		weekday_result.SetText(obtain_weekday(day_str, month_str, year_str))
	})
	calc_button := weekday_form.GetButton(0)

	form_elements := []tview.Primitive{day_field, month_field, year_field, calc_button}
	form_focus := 0

	weekday_form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			form_focus = (form_focus + 1) % len(form_elements)
			app.SetFocus(form_elements[form_focus])
			return nil
		}
		return event
	})


	// Placing of grid elements and configuration of page
	weekday_layout.AddItem(weekday_description, 0, 0, 1, 3, 0, 0, false)
	weekday_layout.AddItem(weekday_form, 1, 0, 1, 3, 0, 0, true)
	weekday_layout.AddItem(weekday_return, 2, 2, 1, 1, 0, 0, false)
	weekday_layout.AddItem(weekday_result, 2, 0, 1, 2, 0, 0, false)

	weekday_layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			current_focus = (current_focus + 1) % len(grid_elements)
			app.SetFocus(grid_elements[current_focus])
			return nil
		} 
		return event 
	})


	return weekday_layout

}