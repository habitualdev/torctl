package ui

import (
	"github.com/rivo/tview"
	"os"
	"torctl/collection"
)

func GetDetails() Terminal {
	app := tview.NewApplication()
	form := tview.NewForm().AddInputField("TOR Host: ", "", 30, nil, nil).
		AddInputField("TOR Control Port: ", "", 30, nil, nil).
		AddDropDown("Auth Type: ", []string{"password", "cookie"}, 0, nil).
		AddPasswordField("Password: ", "", 30, '*', nil).
		AddButton("Save", func() {
			app.Stop()
		}).
		AddButton("Cancel", func() {
			app.Stop()
			os.Exit(0)
		})
	form.SetBorder(true).SetTitle("Torctl").SetTitleAlign(tview.AlignLeft)
	form.SetBorderPadding(1, 1, 1, 1)
	form.SetButtonsAlign(tview.AlignCenter)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	_, authChoice := form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()

	t := Terminal{
		Auth: collection.Connection{
			Auth: collection.ConnectionAuth{
				AuthHost: form.GetFormItem(0).(*tview.InputField).GetText(),
				AuthPort: form.GetFormItem(1).(*tview.InputField).GetText(),
				AuthType: authChoice,
				AuthPass: form.GetFormItem(3).(*tview.InputField).GetText(),
			},
		},
		TerminalData: nil,
	}
	return t
}
