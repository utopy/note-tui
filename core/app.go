package core

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	Root      *tview.Application
	Container *tview.Flex
	Paths     *PathsInterface
}

func CreateApplication() *App {

	instance := tview.NewApplication()

	container := tview.NewFlex()

	paths := StartPathInterface("test_unotes")

	fmt.Println(paths.ConfigPath)

	app := &App{
		Root:      instance,
		Container: container,
		Paths:     paths,
	}

	paths.Setup()

	app.Root.SetRoot(app.Container, true)

	return app
}

func (app *App) LoadLayout(pagesContainer *Pages) {

	// sidebar := tview.NewBox().SetBorder(true).SetTitle("Sidebar")

	sidebar := tview.NewFlex()

	sidebar.SetDirection(tview.FlexColumnCSS)
	sidebar.SetBorder(true)
	sidebar.AddItem(tview.NewBox(), 0, 1, true)

	sidebarFrame := tview.NewFrame(sidebar)
	sidebarFrame.SetBorders(1, 1, 1, 1, 1, 1)
	sidebarFrame.AddText("Notes", true, tview.AlignLeft, tcell.ColorWhite)

	textArea := tview.NewTextArea()
	textArea.SetPlaceholder("Nota")

	pagesContainer.addPage("textArea", textArea, false)

	app.Container.AddItem(sidebarFrame, 35, 0, false)
	app.Container.AddItem(pagesContainer.Container, 0, 1, true)

}

func (app *App) Run() {

	app.Root.EnableMouse(true)

	err := app.Root.Run()

	if err != nil {
		panic("ERROR RUNNING APP")
	}

}
