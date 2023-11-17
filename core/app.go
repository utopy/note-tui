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
	State     *State
}

func CreateApplication() *App {

	instance := tview.NewApplication()

	container := tview.NewFlex()

	paths := StartPathInterface("test_unotes")

	state := InitState("homepage")

	app := &App{
		Root:      instance,
		Container: container,
		Paths:     paths,
		State:     state,
	}

	paths.Setup()

	app.Root.SetRoot(app.Container, true)

	return app
}

func (app *App) LoadLayout(pagesContainer *Pages) {

	sidebar := tview.NewFlex()

	sidebar.SetDirection(tview.FlexColumnCSS)
	sidebar.SetBorder(true)
	sidebar.AddItem(tview.NewBox(), 0, 1, true)

	sidebarFrame := tview.NewFrame(sidebar)
	sidebarFrame.SetBorders(1, 1, 1, 1, 1, 1)
	sidebarFrame.AddText(fmt.Sprintf("[ %s ] mode", app.State.Mode), true, tview.AlignLeft, tcell.ColorWhite)

	app.Root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:

			//Se l'app è già in view mode e si preme esc di nuovo, vai alla home
			if app.State.ActivePage != "homepage" && app.State.Mode == VIEW_MODE {

				pagesContainer.ChangePage("homePage")

				return event

			}

			app.State.ToggleMode(VIEW_MODE)

			sidebarFrame.Clear().AddText(
				fmt.Sprintf("[ %s ] mode", app.State.Mode),
				true,
				tview.AlignLeft,
				tcell.ColorWhite,
			)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'i':
				if app.State.Mode == INSERT_MODE {
					return event
				}
				pagesContainer.ChangePage("textArea")
				app.Root.SetFocus(pagesContainer.Container)
				return nil
			default:
				return event
			}
		default:
			return event
		}
	})

	textArea := tview.NewTextArea()
	textArea.SetPlaceholder("Nota")

	homepage := tview.NewTextView()

	homepage.SetText("Welcome to your tui companion\nPress 'i' to enter edit mode\nPress 'ESC' to enter view mode")

	pagesContainer.AddPage("homePage", homepage, true)
	pagesContainer.AddPage("textArea", textArea, false)

	app.Container.AddItem(sidebarFrame, 35, 0, false)
	app.Container.AddItem(pagesContainer.Frame, 0, 1, true)

}

func (app *App) Run() {

	app.Root.EnableMouse(true)

	err := app.Root.Run()

	if err != nil {
		panic("ERROR RUNNING APP")
	}

}
