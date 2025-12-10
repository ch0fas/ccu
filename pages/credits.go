package pages

import "github.com/rivo/tview"

const page_credits = "CCU Chofas \nThis app exists as a simple collection of date-related functions that I happen to use quite frequently \nThis is a very opinionated application, I understand if it doesn't appeal to everyone, but I will probably use it a lot \nThis project was also created as a way for me to learn tview. \nCheck out the original repo for more info and documentation"

func Credits_Form(app *tview.Application, pages *tview.Pages) tview.Primitive {
	credits_form := tview.NewForm()

	credits_form.SetBorder(true).SetTitle(" Credits ").SetTitleAlign(tview.AlignCenter)

	credits_form.AddTextView("", page_credits, 0, 0, false, false)
	credits_form.AddButton("GO BACK", func() {
		pages.SwitchToPage("main_page")
	})
	credits_form.AddButton("QUIT", func() {
		app.Stop()
	})

	return credits_form
}
