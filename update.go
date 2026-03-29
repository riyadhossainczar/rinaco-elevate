package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateSetup(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		var validator func(string) bool

		switch m.inputStep {
		case 0:
			validator = isValidName
		case 1:
			validator = isValidAge
		case 2:
			validator = isValidAddictYears
		}

		if !validator(m.input) {
			m.hasError = true
			return m, nil
		}

		m.hasError = false
		m.inputBuf = append(m.inputBuf, strings.TrimSpace(m.input))
		m.input = ""
		m.inputStep++

		if m.inputStep == 3 {
			var age int
			var years float64
			fmt.Sscanf(m.inputBuf[1], "%d", &age)
			fmt.Sscanf(m.inputBuf[2], "%f", &years)
			m.data.Name = m.inputBuf[0]
			m.data.Age = age
			m.data.AddictYears = years
			m.data.SetupDone = true
			m.data.LastOpenedAt = m.now
			saveData(m.data)
			m.screen = screenSetStart
			m.inputStep = 0
			m.inputBuf = nil
		}
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
			m.hasError = false
		}
	case "ctrl+c", "q":
		return m, tea.Quit
	default:
		if len(msg.String()) == 1 {
			m.input += msg.String()
			m.hasError = false
		}
	}
	return m, nil
}

func (m model) updateSetStart(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		var validator func(string) bool

		switch m.inputStep {
		case 0:
			validator = isValidYear
		case 1:
			validator = isValidMonth
		case 2:
			validator = isValidDay
		case 3:
			validator = isValidHour
		case 4:
			validator = isValidMinute
		}

		if !validator(m.input) {
			m.hasError = true
			return m, nil
		}

		m.hasError = false
		m.inputBuf = append(m.inputBuf, strings.TrimSpace(m.input))
		m.input = ""
		m.inputStep++

		if m.inputStep == 5 {
			var yr, mo, dy, hr, mn int
			fmt.Sscanf(m.inputBuf[0], "%d", &yr)
			fmt.Sscanf(m.inputBuf[1], "%d", &mo)
			fmt.Sscanf(m.inputBuf[2], "%d", &dy)
			fmt.Sscanf(m.inputBuf[3], "%d", &hr)
			fmt.Sscanf(m.inputBuf[4], "%d", &mn)
			t := time.Date(yr, time.Month(mo), dy, hr, mn, 0, 0, bdTime)

			if len(m.data.Sessions) > 0 {
				last := &m.data.Sessions[len(m.data.Sessions)-1]
				if last.EndedAt.IsZero() {
					last.EndedAt = t
					last.EndedByUs = false
				}
			}
			m.data.Sessions = append(m.data.Sessions, Session{StartedAt: t})
			m.data.LastOpenedAt = m.now
			saveData(m.data)
			m.screen = screenHome
			m.inputStep = 0
			m.inputBuf = nil
		}
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
			m.hasError = false
		}
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.screen = screenHome
		m.input = ""
		m.inputStep = 0
		m.inputBuf = nil
		m.hasError = false
	default:
		if len(msg.String()) == 1 {
			m.input += msg.String()
			m.hasError = false
		}
	}
	return m, nil
}

func (m model) updateHome(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		m.data.LastOpenedAt = m.now
		saveData(m.data)
		return m, tea.Quit
	case "s":
		m.screen = screenSetStart
		m.inputStep = 0
		m.inputBuf = nil
		m.input = ""
		m.hasError = false
	case "r":
		m.screen = screenRelapse
	case "h":
		m.screen = screenHistory
		m.cursor = 0
	case "p":
		m.screen = screenPlan
		m.planTab = 0
		m.cursor = 0
		m.open = -1
	case "t":
		m.screen = screenStats
	case "j":
		m.screen = screenRelapseTrigger
		m.triggerCursor = 0
		m.commandInput = ""
		m.commandError = ""
	case "e":
		m.screen = screenEditProfile
		m.inputStep = 0
		m.inputBuf = nil
		m.input = ""
	}
	return m, nil
}

func (m model) updateStats(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.screen = screenHome
	case "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m model) updateRelapseTrigger(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.screen = screenHome
	case "ctrl+c":
		return m, tea.Quit
	case "j", "down":
		if m.triggerCursor < len(m.data.Sessions)-1 {
			m.triggerCursor++
		}
		m.commandError = ""
	case "k", "up":
		if m.triggerCursor > 0 {
			m.triggerCursor--
		}
		m.commandError = ""
	default:
		if len(msg.String()) == 1 {
			m.commandInput += msg.String()
			m.commandError = ""
		} else if msg.String() == "backspace" {
			if len(m.commandInput) > 0 {
				m.commandInput = m.commandInput[:len(m.commandInput)-1]
				m.commandError = ""
			}
		} else if msg.String() == "enter" {
			cmd := strings.TrimSpace(m.commandInput)
			if strings.HasPrefix(cmd, "v ") {
				idStr := strings.TrimSpace(cmd[2:])
				id, err := strconv.Atoi(idStr)
				if err != nil || id < 1 || id > len(m.data.Sessions) {
					m.commandError = "error"
					m.commandInput = ""
					return m, nil
				}
				m.screen = screenTriggerView
				m.viewingTriggerId = id - 1
				m.editingTrigger = false
				m.tempTriggerReason = m.data.Sessions[m.viewingTriggerId].Reason
				m.commandInput = ""
				m.commandError = ""
			} else if cmd != "" {
				m.commandError = "error"
				m.commandInput = ""
			}
		}
	}
	return m, nil
}

func (m model) updateTriggerView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.screen = screenRelapseTrigger
		m.editingTrigger = false
		m.commandInput = ""
		m.commandError = ""
	case "ctrl+c":
		return m, tea.Quit
	case "c":
		m.editingTrigger = true
		m.tempTriggerReason = m.data.Sessions[m.viewingTriggerId].Reason
	}
	return m, nil
}

func (m model) updateTriggerEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.data.Sessions[m.viewingTriggerId].Reason = m.tempTriggerReason
		saveData(m.data)
		m.editingTrigger = false
	case "backspace":
		if len(m.tempTriggerReason) > 0 {
			m.tempTriggerReason = m.tempTriggerReason[:len(m.tempTriggerReason)-1]
		}
	case "esc":
		m.editingTrigger = false
	case "ctrl+c":
		return m, tea.Quit
	default:
		if len(msg.String()) == 1 {
			m.tempTriggerReason += msg.String()
		}
	}
	return m, nil
}

func (m model) updateRelapse(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "Y":
		m.screen = screenRelapseNote
		m.relapseNote = ""
	case "n", "N", "esc":
		m.screen = screenHome
	case "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m model) updateRelapseNote(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		now := m.now
		if len(m.data.Sessions) > 0 {
			last := &m.data.Sessions[len(m.data.Sessions)-1]
			if last.EndedAt.IsZero() {
				last.EndedAt = now
				last.EndedByUs = true
				last.Reason = m.relapseNote
			}
		}
		m.data.Sessions = append(m.data.Sessions, Session{StartedAt: now})
		m.data.LastOpenedAt = now
		saveData(m.data)
		m.screen = screenHome
	case "esc":
		m.screen = screenRelapse
	case "backspace":
		if len(m.relapseNote) > 0 {
			m.relapseNote = m.relapseNote[:len(m.relapseNote)-1]
		}
	default:
		if len(msg.String()) == 1 {
			m.relapseNote += msg.String()
		}
	}
	return m, nil
}

func (m model) updateHistory(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.screen = screenHome
	case "ctrl+c":
		return m, tea.Quit
	case "j", "down":
		if m.cursor < len(m.data.Sessions)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	}
	return m, nil
}

func (m model) updatePlan(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	items := planSections[m.planTab]
	switch msg.String() {
	case "q", "esc":
		m.screen = screenHome
	case "ctrl+c":
		return m, tea.Quit
	case "tab", "l", "right":
		m.planTab = (m.planTab + 1) % len(planTabs)
		m.cursor = 0
		m.open = -1
	case "shift+tab", "h", "left":
		m.planTab = (m.planTab - 1 + len(planTabs)) % len(planTabs)
		m.cursor = 0
		m.open = -1
	case "j", "down":
		if m.cursor < len(items)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "enter", " ":
		if m.open == m.cursor {
			m.open = -1
		} else {
			m.open = m.cursor
		}
	}
	return m, nil
}

func (m model) updateCheckIn(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "Y":
		m.data.LastOpenedAt = m.now
		saveData(m.data)
		m.screen = screenHome
	case "n", "N":
		if len(m.data.Sessions) > 0 {
			last := &m.data.Sessions[len(m.data.Sessions)-1]
			if last.EndedAt.IsZero() {
				last.EndedAt = m.now
				last.EndedByUs = true
			}
		}
		saveData(m.data)
		m.screen = screenSetStart
		m.inputStep = 0
		m.inputBuf = nil
		m.input = ""
		m.hasError = false
	case "ctrl+c":
		return m, tea.Quit
	}
	return m, nil
}

func (m model) updateEditProfile(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if m.inputBuf == nil {
			m.inputBuf = make([]string, 0)
		}

		var validator func(string) bool

		switch m.inputStep {
		case 0:
			validator = isValidName
		case 1:
			validator = isValidAge
		case 2:
			validator = isValidAddictYears
		}

		if !validator(m.input) {
			m.hasError = true
			return m, nil
		}

		m.hasError = false
		m.inputBuf = append(m.inputBuf, strings.TrimSpace(m.input))
		m.input = ""
		m.inputStep++

		if m.inputStep == 3 {
			var age int
			var years float64
			fmt.Sscanf(m.inputBuf[1], "%d", &age)
			fmt.Sscanf(m.inputBuf[2], "%f", &years)
			m.data.Name = m.inputBuf[0]
			m.data.Age = age
			m.data.AddictYears = years
			saveData(m.data)
			m.screen = screenHome
			m.inputStep = 0
			m.inputBuf = nil
		}
	case "backspace":
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
			m.hasError = false
		}
	case "ctrl+c", "esc":
		m.screen = screenHome
		m.input = ""
		m.inputStep = 0
		m.inputBuf = nil
	default:
		if len(msg.String()) == 1 {
			m.input += msg.String()
			m.hasError = false
		}
	}
	return m, nil
}
