package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	switch m.screen {
	case screenSetup:
		return m.viewSetup()
	case screenHome:
		return m.viewHome()
	case screenStats:
		return m.viewStats()
	case screenRelapseTrigger:
		return m.viewRelapseTrigger()
	case screenSetStart:
		return m.viewSetStart()
	case screenRelapse:
		return m.viewRelapse()
	case screenRelapseNote:
		return m.viewRelapseNote()
	case screenHistory:
		return m.viewHistory()
	case screenPlan:
		return m.viewPlan()
	case screenCheckIn:
		return m.viewCheckIn()
	case screenEditProfile:
		return m.viewEditProfile()
	case screenTriggerView:
		if m.editingTrigger {
			return m.viewTriggerEdit()
		}
		return m.viewTriggerView()
	}
	return ""
}

func (m model) viewSetup() string {
	prompts := []string{
		"Your name: ",
		"Your age: ",
		"Years of addiction (e.g. 5.5): ",
	}

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(white.Render(" Rinaco Elevate") + "  " + dimStyle.Render(version) + "\n")
	b.WriteString(dimStyle.Render("  First time setup") + "\n\n")

	for i, p := range prompts {
		if i < m.inputStep {
			b.WriteString(good.Render("  ✓ ") + dimStyle.Render(p+m.inputBuf[i]) + "\n")
		} else if i == m.inputStep {
			errorStr := ""
			if m.hasError {
				errorStr = "  " + redStyle.Render("✗")
			}
			b.WriteString("  " + accent.Render(p) + white.Render(m.input) + accent.Render("▌") + errorStr + "\n")
		} else {
			b.WriteString(dimStyle.Render("  "+p) + "\n")
		}
	}
	b.WriteString("\n" + dimStyle.Render("  press enter after each\n"))
	return b.String()
}

func (m model) viewSetStart() string {
	stepLabels := []string{
		"Year (e.g. 2026): ",
		"Month (1-12): ",
		"Day: ",
		"Hour (0-23): ",
		"Minute (0-59): ",
	}
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(white.Render("  Set start time") + "\n")
	b.WriteString(dimStyle.Render("  When did you start this streak?") + "\n")
	b.WriteString(dimStyle.Render("  (BST UTC+6)") + "\n\n")

	for i, lbl := range stepLabels {
		if i < m.inputStep {
			b.WriteString(good.Render("  ✓ ") + dimStyle.Render(lbl+m.inputBuf[i]) + "\n")
		} else if i == m.inputStep {
			errorStr := ""
			if m.hasError {
				errorStr = "  " + redStyle.Render("✗")
			}
			b.WriteString("  " + accent.Render(lbl) + white.Render(m.input) + accent.Render("▌") + errorStr + "\n")
		} else {
			b.WriteString(dimStyle.Render("  "+lbl) + "\n")
		}
	}
	b.WriteString("\n" + dimStyle.Render("  esc to cancel\n"))
	return b.String()
}

func (m model) viewHome() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(white.Render("  Rinaco Elevate") + "  " + dimStyle.Render(version) + "\n")
	b.WriteString(dimStyle.Render(fmt.Sprintf("  %s  |  %d y/o  |  %.1fy addiction", m.data.Name, m.data.Age, m.data.AddictYears)) + "\n\n")

	start := currentStart(m.data)
	if start.IsZero() {
		b.WriteString(warn.Render("  no active streak\n\n"))
	} else {
		dur := m.now.Sub(start)
		totalMins := int64(dur.Minutes())

		level, current, next := getLevelStats(totalMins)
		progress := float64(totalMins-current) / float64(next-current)
		if progress < 0 {
			progress = 0
		}
		progressBar := ""
		barLen := 20
		filledLen := int(progress * float64(barLen))
		for i := 0; i < barLen; i++ {
			if i < filledLen {
				progressBar += "█"
			} else {
				progressBar += "░"
			}
		}

		b.WriteString(good.Render("  ◉ active") + "\n\n")
		b.WriteString(white.Render(fmt.Sprintf("  %s", formatDuration(dur))) + "\n")
		b.WriteString(dimStyle.Render(fmt.Sprintf("  %d minutes  |  since %s", totalMins, start.In(bdTime).Format("02 Jan, 15:04"))) + "\n\n")
		b.WriteString(levelStyle.Render(fmt.Sprintf("  ▲ LEVEL %d", level)) + dimStyle.Render(fmt.Sprintf("  [%d → %d min]", current, next)) + "\n")

		b.WriteString("  " + accent.Render(progressBar) + "\n\n")
	}

	b.WriteString(dimStyle.Render("  ───────────────────────────────") + "\n")

	b.WriteString("  " + muted.Render("s") + dimStyle.Render(" set start       ") + muted.Render("r") + dimStyle.Render(" relapsed") + "\n")
	b.WriteString("  " + muted.Render("h") + dimStyle.Render(" history         ") + muted.Render("p") + dimStyle.Render(" plan") + "\n")
	b.WriteString("  " + muted.Render("t") + dimStyle.Render(" stats           ") + muted.Render("j") + dimStyle.Render(" triggers") + "\n")
	b.WriteString("  " + muted.Render("e") + dimStyle.Render(" edit profile    ") + muted.Render("q") + dimStyle.Render(" quit") + "\n")

	return b.String()
}

func (m model) viewStats() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  Statistics") + "\n\n")

	if len(m.data.Sessions) == 0 {
		b.WriteString(dimStyle.Render("  no data yet\n\n"))
		b.WriteString(dimStyle.Render("  esc to go back\n"))
		return b.String()
	}

	longest := getLongestStreak(m.data.Sessions)
	b.WriteString(accent.Render("  ▸ Longest Streak") + "\n")
	b.WriteString(good.Render(fmt.Sprintf("    %s", formatDuration(longest))) + "\n\n")

	totalDays := getTotalCleanDays(m.data.Sessions)
	b.WriteString(accent.Render("  ▸ Current Streak") + "\n")

	b.WriteString(good.Render(fmt.Sprintf("    %d days", totalDays)) + "\n\n")

	relapses := getRelapseCount(m.data.Sessions)

	b.WriteString(redStyle.Render(fmt.Sprintf("  ✗ Relapses: %d", relapses)) + "\n\n")

	hour, count := getDangerousTimeWindow(m.data.Sessions)
	if count > 0 {
		b.WriteString(warn.Render(fmt.Sprintf("  ⚠ Danger Zone: %02d:00 (%d times)", hour, count)) + "\n\n")
	}

	currentMonth, previousMonth := getComparisonStats(m.data.Sessions, m.now)
	change := currentMonth - previousMonth
	changeStr := fmt.Sprintf("%+d", change)
	changeColor := good
	if change < 0 {
		changeColor = redStyle
	}

	b.WriteString(accent.Render("  ▸ Month Comparison") + "\n")

	b.WriteString(dimStyle.Render(fmt.Sprintf("    This month: %d  ", currentMonth)))

	b.WriteString(changeColor.Render(changeStr) + "\n\n")

	b.WriteString(dimStyle.Render("  esc to go back\n"))
	return b.String()
}

func (m model) viewRelapseTrigger() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  Relapse Triggers") + "\n\n")

	if len(m.data.Sessions) == 0 {
		b.WriteString(dimStyle.Render("  no relapses recorded\n\n"))
		b.WriteString(dimStyle.Render("  esc to go back\n"))
		return b.String()
	}

	relapseCount := 0
	for i, s := range m.data.Sessions {
		if !s.EndedByUs {
			continue
		}

		prefix := "  "
		if relapseCount == m.triggerCursor {
			prefix = accent.Render("  > ")
		} else {
			prefix = "    "
		}

		id := fmt.Sprintf("[%d]", i+1)
		date := s.StartedAt.In(bdTime).Format("02 Jan 15:04")
		reason := s.Reason
		if reason == "" {
			reason = dimStyle.Render("(no reason)")
		}

		b.WriteString(prefix + white.Render(id) + "  " + date + "  " + reason + "\n")
		relapseCount++
	}

	errorDisplay := ""
	if m.commandError != "" {
		errorDisplay = "  " + redStyle.Render(m.commandError) + "\n"
	}

	b.WriteString("\n")
	b.WriteString(accent.Render("  command:") + " " + white.Render(m.commandInput) + accent.Render("▌") + "\n")
	b.WriteString(errorDisplay)
	b.WriteString("\n" + dimStyle.Render("  j/k navigate   v [id] view   esc back\n"))
	return b.String()
}

func (m model) viewTriggerView() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  Relapse Details") + "\n\n")

	s := m.data.Sessions[m.viewingTriggerId]
	date := s.StartedAt.In(bdTime).Format("02 January 2006, 15:04")

	b.WriteString(accent.Render("  ID:") + dimStyle.Render(fmt.Sprintf("  [%d]", m.viewingTriggerId+1)) + "\n")

	b.WriteString(accent.Render("  Date:") + dimStyle.Render(fmt.Sprintf("  %s", date)) + "\n\n")

	b.WriteString(accent.Render("  Reason:") + "\n")

	if s.Reason == "" {
		b.WriteString(dimStyle.Render("    (empty)") + "\n")
	} else {
		for _, line := range strings.Split(s.Reason, "\n") {

			b.WriteString("    " + dimStyle.Render(line) + "\n")
		}
	}

	b.WriteString("\n" + dimStyle.Render("  c edit   esc back") + "\n")
	return b.String()
}

func (m model) viewTriggerEdit() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  Edit Reason") + "\n\n")

	b.WriteString(accent.Render("  Write reason:") + "\n\n")
	b.WriteString("  " + white.Render(m.tempTriggerReason) + accent.Render("▌") + "\n\n")
	b.WriteString(dimStyle.Render("  enter save   esc cancel") + "\n")
	return b.String()
}

func (m model) viewRelapse() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(redStyle.Render("  relapsed?") + "\n\n")

	b.WriteString(dimStyle.Render("  your streak will end now.") + "\n\n")
	b.WriteString(white.Render("  y") + dimStyle.Render("  yes, relapsed") + "\n")
	b.WriteString(white.Render("  n") + dimStyle.Render("  no, go back") + "\n")

	return b.String()
}

func (m model) viewRelapseNote() string {
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(white.Render("  What triggered it?") + "\n\n")

	b.WriteString(dimStyle.Render("  Write why you relapsed:") + "\n\n")
	b.WriteString("  " + white.Render(m.relapseNote) + accent.Render("▌") + "\n\n")
	b.WriteString(dimStyle.Render("  enter save   esc skip") + "\n")
	return b.String()
}

func (m model) viewHistory() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  History") + "\n\n")

	if len(m.data.Sessions) == 0 {
		b.WriteString(dimStyle.Render("  no sessions yet\n\n"))
		b.WriteString(dimStyle.Render("  esc to go back\n"))
		return b.String()
	}

	for i, s := range m.data.Sessions {
		prefix := "  "
		if i == m.cursor {
			prefix = accent.Render("  > ")
		} else {
			prefix = "    "
		}

		num := fmt.Sprintf("#%d", i+1)
		start := s.StartedAt.In(bdTime).Format("02 Jan 15:04")

		if s.EndedAt.IsZero() {
			dur := m.now.Sub(s.StartedAt)
			b.WriteString(prefix + good.Render(num) + "  " + white.Render(start) + dimStyle.Render(" → now  ") + good.Render(formatDuration(dur)) + "\n")
		} else {
			dur := s.EndedAt.Sub(s.StartedAt)
			end := s.EndedAt.In(bdTime).Format("02 Jan 15:04")
			tag := dimStyle.Render("ended")
			if s.EndedByUs {
				tag = redStyle.Render("relapsed")
			}
			b.WriteString(prefix + muted.Render(num) + "  " + dimStyle.Render(start) + dimStyle.Render(" → ") + dimStyle.Render(end) + "  " + tag + "  " + muted.Render(formatDuration(dur)) + "\n")
		}
	}

	b.WriteString("\n" + dimStyle.Render("  j/k navigate   esc back\n"))
	return b.String()
}

func (m model) viewPlan() string {
	var b strings.Builder
	b.WriteString("\n" + white.Render("  Plan") + "\n\n")

	b.WriteString("  ")
	for i, t := range planTabs {
		if i == m.planTab {
			b.WriteString(activeTab.Render("[" + t + "]"))
		} else {
			b.WriteString(inactTab.Render(" " + t + " "))
		}
		if i < len(planTabs)-1 {
			b.WriteString(dimStyle.Render("  "))
		}
	}
	b.WriteString("\n\n")

	items := planSections[m.planTab]
	for i, item := range items {
		prefix := "    "
		if i == m.cursor {
			prefix = accent.Render("  > ")
		}

		titleLine := white.Render(item.title)
		if item.sub != "" {
			titleLine += "  " + muted.Render(item.sub)
		}
		ind := dimStyle.Render(" ▸")
		if m.open == i {
			ind = dimStyle.Render(" ▾")
		}

		b.WriteString(prefix + titleLine + ind + "\n")

		if m.open == i && item.detail != "" {
			for _, line := range strings.Split(item.detail, "\n") {
				if strings.HasPrefix(line, "→") {
					b.WriteString("      " + warn.Render(line) + "\n")
				} else {
					b.WriteString("      " + italic.Render(line) + "\n")
				}
			}
			b.WriteString("\n")
		}
	}

	b.WriteString("\n" + dimStyle.Render("  j/k nav   enter expand   tab switch   esc back\n"))
	return b.String()
}

func (m model) viewCheckIn() string {
	away := m.now.Sub(m.data.LastOpenedAt)
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(warn.Render(fmt.Sprintf("  away for %s", formatDuration(away))) + "\n\n")
	b.WriteString(white.Render("  still clean?") + "\n\n")
	b.WriteString(white.Render("  y") + dimStyle.Render("  yes\n"))
	b.WriteString(white.Render("  n") + dimStyle.Render("  no, relapsed\n"))
	return b.String()
}

func (m model) viewEditProfile() string {
	prompts := []string{
		"Your name: ",
		"Your age: ",
		"Years of addiction: ",
	}

	var b strings.Builder
	b.WriteString("\n" + white.Render("  Edit Profile") + "\n\n")

	for i, p := range prompts {
		if i < m.inputStep {
			b.WriteString(good.Render("  ✓ ") + dimStyle.Render(p+m.inputBuf[i]) + "\n")
		} else if i == m.inputStep {
			errorStr := ""
			if m.hasError {
				errorStr = "  " + redStyle.Render("✗")
			}
			b.WriteString("  " + accent.Render(p) + white.Render(m.input) + accent.Render("▌") + errorStr + "\n")
		} else {
			b.WriteString(dimStyle.Render("  "+p) + "\n")
		}
	}
	b.WriteString("\n" + dimStyle.Render("  enter after each   esc cancel\n"))
	return b.String()
}
