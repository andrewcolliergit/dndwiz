package shell

import (
	"dndwiz/tui/style"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ContentView struct {
	Viewport viewport.Model
}

func NewContentView(width, height int) ContentView {
	vp := viewport.New(width, height)
	return ContentView{Viewport: vp}
}

func (cv *ContentView) Update(msg tea.Msg) (ContentView, tea.Cmd) {
	var cmd tea.Cmd
	cv.Viewport, cmd = cv.Viewport.Update(msg)
	return *cv, cmd
}

func (cv ContentView) View() string {
	return style.ContentBorder.Render(cv.Viewport.View())
}

func (cv *ContentView) SetSize(width, height int) {
	cv.Viewport.Width = width
	cv.Viewport.Height = height
}

func (cv *ContentView) SetContent(content string) {
	cv.Viewport.SetContent(content)
}
