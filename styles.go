package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	dimStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	white      = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	activeTab  = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	inactTab   = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
	accent     = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	warn       = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
	good       = lipgloss.NewStyle().Foreground(lipgloss.Color("71"))
	muted      = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	italic     = lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Italic(true)
	redStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	levelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("135")).Bold(true)
)
