package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	messages    []string
	ClearInput  func()
	SetMessages func([]string)
	OnSubmit    func(string)
}

func (ui *UI) Start() {
	app := tview.NewApplication()
	// Input field for typing messages
	inputField := tview.NewInputField().
		SetLabel("Type your message: ").
		SetAcceptanceFunc(tview.InputFieldMaxLength(100))
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			ui.OnSubmit(inputField.GetText())
			inputField.SetText("")
		}
	})

	// Text view for displaying chat messages
	chatTextView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
		SetText("Welcome to the Chat\n")

	// Flex layout to arrange components vertically
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(chatTextView, 0, 1, false).
		AddItem(inputField, 3, 1, true)

	// Set the root component of the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
