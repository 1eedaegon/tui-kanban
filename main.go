package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

const (
	todo status = iota
	inprogress
	done
)

// TASK
type Task struct {
	status status
	title  string
	desc   string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.desc
}

// Model
type Model struct {
	lists []list.Model
	err   error
}

func New() *Model {
	return &Model{}
}

// TODO: Call this on tea.WindowSizeMsg
func (m *Model) initList(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	m.lists[todo].Title = "To do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "coding practice", desc: "oauth practice"},
		Task{status: todo, title: "struct odroid", desc: "odroid -n2"},
		Task{status: todo, title: "fold laundry", desc: "shirts"},
	})
	m.lists[inprogress].Title = "In progress"
	m.lists[inprogress].SetItems([]list.Item{
		Task{status: inprogress, title: "coding practice", desc: "oauth practice"},
		Task{status: inprogress, title: "struct odroid", desc: "odroid -n2"},
		Task{status: inprogress, title: "fold laundry", desc: "shirts"},
	})
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "coding practice", desc: "oauth practice"},
		Task{status: done, title: "struct odroid", desc: "odroid -n2"},
		Task{status: done, title: "fold laundry", desc: "shirts"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
	}
	var cmd tea.Cmd
	m.lists, cmd = m.lists.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.lists.View()
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}
}
