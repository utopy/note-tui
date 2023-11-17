package core

import "github.com/rivo/tview"

type Pages struct {
	Container *tview.Pages
	List      []string
}

func InitPages() *Pages {

	container := tview.NewPages()

	pages := &Pages{
		Container: container,
		List:      make([]string, 0),
	}

	pages.Container.AddAndSwitchToPage("test page", tview.NewBox().SetTitle(" test page ").SetBorder(true), true)

	return pages
}

func (pages *Pages) addPage(pageName string, item tview.Primitive, hasFocus bool) {

	pages.Container.AddPage(pageName, item, true, hasFocus)

	pages.List = append(pages.List, pageName)

}
