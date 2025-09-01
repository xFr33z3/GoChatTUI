package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	messages := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	input := tview.NewInputField().
		SetLabel("Message: ").
		SetFieldWidth(0)

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			offsetRow, offsetColumn := messages.GetScrollOffset()
			messages.ScrollTo(offsetRow-1, offsetColumn)
		case tcell.KeyDown:
			offsetRow, offsetColumn := messages.GetScrollOffset()
			messages.ScrollTo(offsetRow+1, offsetColumn)
		}
		return event
	})

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(messages, 0, 1, false).
		AddItem(input, 1, 0, true)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Fprintf(messages, "[red]Disconnected from server\n")
				return
			}
			messages.ScrollToEnd()
			fmt.Fprintf(messages, "%s", string(buf[:n]))
		}
	}()

	input.SetDoneFunc(func(key tcell.Key) {
		text := input.GetText()
		if text != "" {
			conn.Write([]byte(text + "\n"))
			input.SetText("")
		}
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
