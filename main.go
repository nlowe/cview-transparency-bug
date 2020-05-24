package main

import (
	"os"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"gitlab.com/tslocum/cview"
)

var (
	cviewCmd = &cobra.Command{
		Use:  "cview",
		Args: cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			app := cview.NewApplication()

			// Returns a new primitive which puts the provided primitive in the center and
			// sets its size to the given width and height.
			modal := func(p cview.Primitive, width, height int) cview.Primitive {
				return cview.NewFlex().
					AddItem(nil, 0, 1, false).
					AddItem(cview.NewFlex().SetDirection(cview.FlexRow).
						AddItem(nil, 0, 1, false).
						AddItem(p, height, 1, false).
						AddItem(nil, 0, 1, false), width, 1, false).
					AddItem(nil, 0, 1, false)
			}

			background := cview.NewTextView().
				SetTextColor(tcell.ColorBlue).
				SetText(strings.Repeat("background ", 1000))

			box := cview.NewBox().
				SetBorder(true).
				SetTitle("Centered Box")

			pages := cview.NewPages().
				AddPage("background", background, true, true).
				AddPage("modal", modal(box, 40, 10), true, true)

			return app.SetRoot(pages, true).Run()
		},
	}

	tviewCmd = &cobra.Command{
		Use:  "tview",
		Args: cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			app := tview.NewApplication()

			// Returns a new primitive which puts the provided primitive in the center and
			// sets its size to the given width and height.
			modal := func(p tview.Primitive, width, height int) tview.Primitive {
				return tview.NewFlex().
					AddItem(nil, 0, 1, false).
					AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
						AddItem(nil, 0, 1, false).
						AddItem(p, height, 1, false).
						AddItem(nil, 0, 1, false), width, 1, false).
					AddItem(nil, 0, 1, false)
			}

			background := tview.NewTextView().
				SetTextColor(tcell.ColorBlue).
				SetText(strings.Repeat("background ", 1000))

			box := tview.NewBox().
				SetBorder(true).
				SetTitle("Centered Box")

			pages := tview.NewPages().
				AddPage("background", background, true, true).
				AddPage("modal", modal(box, 40, 10), true, true)

			return app.SetRoot(pages, true).Run()
		},
	}
)

func main() {
	root := &cobra.Command{}
	root.AddCommand(cviewCmd, tviewCmd)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
