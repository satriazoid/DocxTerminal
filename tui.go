package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	screenMenu screen = iota
	screenList
	screenView
	screenAdd
	screenEditPick
	screenDelete
	screenConfirm
)

type editorDoneMsg struct{ err error }

const logo = `
  ____             _____                    _             _
 |  _ \  ___   ___|_   _|__ _ __ _ __ ___ (_)_ __   __ _| |
 | | | |/ _ \ / __| | |/ _ \ '__| '_ \ _ \| | '_ \ / _' | |
 | |_| | (_) | (__  | |  __/ |  | | | | | | | | | | (_| | |
 |____/ \___/ \___| |_|\___|_|  |_| |_| |_|_|_| |_|\__,_|_|`

var (
	cAccent = lipgloss.Color("81")
	cMuted  = lipgloss.Color("240")
	cText   = lipgloss.Color("252")
	cSelBg  = lipgloss.Color("236")
	cWarn   = lipgloss.Color("203")
	cOk     = lipgloss.Color("114")

	stTitle = lipgloss.NewStyle().Bold(true).Foreground(cAccent)
	stMuted = lipgloss.NewStyle().Foreground(cMuted)
	stSel   = lipgloss.NewStyle().Foreground(cAccent).Background(cSelBg).Bold(true)
	stNorm  = lipgloss.NewStyle().Foreground(cText)
	stErr   = lipgloss.NewStyle().Foreground(cWarn).Bold(true)
	stOk    = lipgloss.NewStyle().Foreground(cOk)
	stBox   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(cAccent).Padding(0, 1)
	stHelp  = lipgloss.NewStyle().Foreground(cMuted)
)

type tui struct {
	store    *Store
	screen   screen
	cursor   int
	menu     []string
	docs     []string
	viewName string
	content  string
	status   string
	input    textinput.Model
	vp       viewport.Model
	w, h     int
	confirm  string
	pending  string
}

func newTUI(s *Store) tui {
	ti := textinput.New()
	ti.Placeholder = "name (git, pnpm, ...)"
	ti.CharLimit = 64
	ti.Width = 40
	docs, _ := s.List()
	return tui{
		store: s,
		menu:  []string{"List", "Add", "Edit", "Delete", "Quit"},
		docs:  docs,
		input: ti,
		vp:    viewport.New(80, 20),
	}
}

func (m tui) Init() tea.Cmd { return nil }

func (m tui) refresh() tui {
	docs, err := m.store.List()
	if err != nil {
		m.status = err.Error()
		return m
	}
	m.docs = docs
	if m.cursor >= len(m.docs) && m.cursor > 0 {
		m.cursor = len(m.docs) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	return m
}

func (m tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = msg.Width, msg.Height
		inner := max(10, m.h-12)
		m.vp.Width = max(20, m.w-6)
		m.vp.Height = inner
		if m.content != "" {
			m.vp.SetContent(renderMarkdown(m.content, m.vp.Width))
		}
		return m, nil

	case editorDoneMsg:
		if msg.err != nil {
			m.status = msg.err.Error()
		} else {
			m.status = "saved"
			m = m.refresh()
			if m.viewName != "" {
				if body, err := m.store.Read(m.viewName); err == nil {
					m.content = body
					m.vp.SetContent(renderMarkdown(body, m.vp.Width))
					m.screen = screenView
				}
			}
		}
		return m, nil

	case tea.KeyMsg:
		if m.screen == screenAdd {
			return m.updateAdd(msg)
		}
		if m.screen == screenConfirm {
			return m.updateConfirm(msg)
		}
		if m.screen == screenView {
			return m.updateView(msg)
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			if m.screen == screenMenu {
				return m, tea.Quit
			}
			m.screen = screenMenu
			m.cursor = 0
			m.status = ""
			return m, nil
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			m.cursor++
			lim := m.listLen() - 1
			if m.cursor > lim {
				m.cursor = lim
			}
			if m.cursor < 0 {
				m.cursor = 0
			}
		case "enter":
			return m.enter()
		}
	}
	return m, nil
}

func (m tui) listLen() int {
	switch m.screen {
	case screenMenu:
		return len(m.menu)
	case screenList, screenEditPick, screenDelete:
		return len(m.docs)
	}
	return 1
}

func (m tui) enter() (tea.Model, tea.Cmd) {
	switch m.screen {
	case screenMenu:
		switch m.cursor {
		case 0:
			m = m.refresh()
			m.screen = screenList
			m.cursor = 0
			m.status = ""
		case 1:
			m.screen = screenAdd
			m.input.SetValue("")
			m.input.Focus()
			m.status = ""
			return m, textinput.Blink
		case 2:
			m = m.refresh()
			m.screen = screenEditPick
			m.cursor = 0
			m.status = ""
		case 3:
			m = m.refresh()
			m.screen = screenDelete
			m.cursor = 0
			m.status = ""
		case 4:
			return m, tea.Quit
		}
	case screenList:
		if len(m.docs) == 0 {
			m.status = "empty"
			return m, nil
		}
		name := m.docs[m.cursor]
		body, err := m.store.Read(name)
		if err != nil {
			m.status = err.Error()
			return m, nil
		}
		m.viewName = name
		m.content = body
		m.vp.SetContent(renderMarkdown(body, m.vp.Width))
		m.vp.GotoTop()
		m.screen = screenView
	case screenEditPick:
		if len(m.docs) == 0 {
			m.status = "empty"
			return m, nil
		}
		return m.openEditor(m.docs[m.cursor])
	case screenDelete:
		if len(m.docs) == 0 {
			m.status = "empty"
			return m, nil
		}
		m.pending = m.docs[m.cursor]
		m.confirm = "delete " + m.pending + "?"
		m.screen = screenConfirm
	}
	return m, nil
}

func (m tui) openEditor(name string) (tea.Model, tea.Cmd) {
	p, err := m.store.FilePath(name)
	if err != nil {
		m.status = err.Error()
		return m, nil
	}
	m.viewName = name
	cmd := editorCmd(p)
	return m, tea.ExecProcess(cmd, func(err error) tea.Msg {
		return editorDoneMsg{err: err}
	})
}

func (m tui) updateAdd(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.screen = screenMenu
		m.input.Blur()
		return m, nil
	case "enter":
		name := strings.TrimSpace(m.input.Value())
		if name == "" {
			m.status = "name required"
			return m, nil
		}
		if err := m.store.Create(name); err != nil {
			m.status = err.Error()
			return m, nil
		}
		m.input.Blur()
		return m.openEditor(name)
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m tui) updateConfirm(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "Y", "enter":
		if err := m.store.Delete(m.pending); err != nil {
			m.status = err.Error()
		} else {
			m.status = "deleted " + m.pending
		}
		m.pending = ""
		m = m.refresh()
		m.screen = screenDelete
		return m, nil
	case "n", "N", "esc", "q":
		m.pending = ""
		m.screen = screenDelete
		m.status = "cancelled"
		return m, nil
	}
	return m, nil
}

func renderMarkdown(src string, width int) string {
	if width < 20 {
		width = 20
	}
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return src
	}
	out, err := r.Render(src)
	if err != nil {
		return src
	}
	return out
}

func (m tui) updateView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.screen = screenList
		return m, nil
	case "e":
		return m.openEditor(m.viewName)
	case "ctrl+c":
		return m, tea.Quit
	}
	var cmd tea.Cmd
	m.vp, cmd = m.vp.Update(msg)
	return m, cmd
}

func (m tui) View() string {
	if m.w == 0 {
		return "loading..."
	}
	banner := stTitle.Render(logo)
	header := banner + "\n" + stMuted.Render("  markdown cheat sheets  ·  ~/.docxterminal")
	body := ""
	switch m.screen {
	case screenMenu:
		body = m.renderChoices(m.menu)
	case screenList:
		body = stMuted.Render("list") + "\n" + m.renderDocs()
	case screenEditPick:
		body = stMuted.Render("edit — pick") + "\n" + m.renderDocs()
	case screenDelete:
		body = stMuted.Render("delete — pick") + "\n" + m.renderDocs()
	case screenAdd:
		body = stMuted.Render("add") + "\n\n" + m.input.View()
	case screenView:
		body = stMuted.Render(m.viewName+".md") + "\n" + m.vp.View()
	case screenConfirm:
		body = stErr.Render(m.confirm) + "\n\n" + stMuted.Render("y confirm · n cancel")
	}
	status := m.status
	if status == "" {
		status = " "
	}
	help := m.help()
	innerW := max(20, m.w-4)
	box := stBox.Width(innerW).Render(header + "\n\n" + body)
	return box + "\n" + stOk.Render(status) + "\n" + stHelp.Render(help)
}

func (m tui) renderChoices(items []string) string {
	var b strings.Builder
	for i, it := range items {
		cur := "  "
		style := stNorm
		if i == m.cursor {
			cur = "▸ "
			style = stSel
		}
		b.WriteString(style.Render(cur+it) + "\n")
	}
	return b.String()
}

func (m tui) renderDocs() string {
	if len(m.docs) == 0 {
		return stMuted.Render("\n(no documents)")
	}
	return m.renderChoices(m.docs)
}

func (m tui) help() string {
	switch m.screen {
	case screenMenu:
		return "j/k move · enter · q quit"
	case screenList:
		return "j/k move · enter view · q back"
	case screenView:
		return "rendered markdown · e edit · j/k scroll · q back"
	case screenAdd:
		return "enter create+edit · esc back"
	case screenEditPick:
		return "j/k move · enter edit · q back"
	case screenDelete:
		return "j/k move · enter delete · q back"
	case screenConfirm:
		return "y / n"
	}
	return ""
}

func runTUI(s *Store) error {
	p := tea.NewProgram(newTUI(s), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
