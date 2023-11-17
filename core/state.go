package core

type Mode string

type State struct {
	ActivePage string
	Mode       Mode
}

const (
	INSERT_MODE Mode = "insert"
	VIEW_MODE   Mode = "view"
)

func InitState(activePage string) *State {
	return &State{
		ActivePage: activePage,
		Mode:       VIEW_MODE,
	}
}

func (state *State) ToggleMode(mode Mode) {
	state.Mode = mode
}
