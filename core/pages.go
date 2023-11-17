package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Pages struct {
	Container *tview.Pages
	Frame     *tview.Frame
	List      []string
}

func InitPages() *Pages {

	container := tview.NewPages()
	containerFrame := tview.NewFrame(container)
	containerFrame.SetBorders(1, 1, 1, 1, 1, 1)
	containerFrame.AddText("Notes", true, tview.AlignLeft, tcell.ColorWhite)

	pages := &Pages{
		Container: container,
		Frame:     containerFrame,
		List:      make([]string, 0),
	}

	pages.Container.AddAndSwitchToPage("test page", tview.NewBox().SetTitle(" test page ").SetBorder(true), true)

	return pages
}

func (pages *Pages) AddPage(pageName string, item tview.Primitive, hasFocus bool) {

	pages.Container.AddPage(pageName, item, true, hasFocus)

	pages.List = append(pages.List, pageName)

}

func (pages *Pages) ChangePage(pageName string) {
	pages.Container.SwitchToPage(pageName)
}
