package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)








func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}





func newModel() model {
	d := loadData()
	now := time.Now().In(bdTime)

	m := model{
		data:     d,
		open:     -1,
		now:      now,
		hasError: false,
	}

	if !d.SetupDone {
		m.screen = screenSetup
		m.inputStep = 0
		return m
	}

	if !d.LastOpenedAt.IsZero() && now.Sub(d.LastOpenedAt) >= checkGap {
		start := currentStart(d)
		if !start.IsZero() {
			m.screen = screenCheckIn
			return m
		}
	}

	m.screen = screenHome
	return m
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		m.now = time.Time(msg).In(bdTime)
		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch m.screen {
		case screenSetup:
			return m.updateSetup(msg)
		case screenHome:
			return m.updateHome(msg)
		case screenStats:
			return m.updateStats(msg)
		case screenRelapseTrigger:
			return m.updateRelapseTrigger(msg)
		case screenSetStart:
			return m.updateSetStart(msg)
		case screenRelapse:
			return m.updateRelapse(msg)
		case screenRelapseNote:
			return m.updateRelapseNote(msg)
		case screenHistory:
			return m.updateHistory(msg)
		case screenPlan:
			return m.updatePlan(msg)
		case screenCheckIn:
			return m.updateCheckIn(msg)
		case screenEditProfile:
			return m.updateEditProfile(msg)
		case screenTriggerView:
			if m.editingTrigger {
				return m.updateTriggerEdit(msg)
			}
			return m.updateTriggerView(msg)
		}
	}
	return m, nil
}
