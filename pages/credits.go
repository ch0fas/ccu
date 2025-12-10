package pages

import "github.com/rivo/tview"

const page_credits = "CCU Chofas \nThis app was made by ch0fas with ❤️ \nThis app is distributed under a GNU GPLv3 License. See original repo for more details."

func Credits_Form(app *tview.Application, pages *tview.Pages) tview.Primitive {
	credits_form := tview.NewForm()

	credits_form.SetBorder(true).SetTitle(" Credits ").SetTitleAlign(tview.AlignCenter)

	credits_form.AddTextView("", page_credits, 0, 0, false, false)
	credits_form.AddButton("GO BACK", func(){
		pages.SwitchToPage("main_page")
	})
	credits_form.AddButton("QUIT", func() {
		app.Stop()
	})



	return credits_form
}