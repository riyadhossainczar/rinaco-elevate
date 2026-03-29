package main

import (
	"time"
)

const (
	version  = "v1.5.19-B3 [FIB] #7a2b9f1"
	dataFile = "rinaco-elevate_data.json"
	checkGap = 48 * time.Hour
)

var bdTime = time.FixedZone("BST", 6*60*60)

type Session struct {
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at,omitempty"`
	EndedByUs bool      `json:"ended_by_user"`
	Reason    string    `json:"reason,omitempty"`
}

type UserData struct {
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	AddictYears  float64   `json:"addict_years"`
	Sessions     []Session `json:"sessions"`
	LastOpenedAt time.Time `json:"last_opened_at"`
	SetupDone    bool      `json:"setup_done"`
}

type screen int

const (
	screenSetup screen = iota
	screenHome
	screenStats
	screenRelapseTrigger
	screenSetStart
	screenRelapse
	screenRelapseNote
	screenHistory
	screenPlan
	screenCheckIn
	screenEditProfile
	screenTriggerView
	screenTriggerEdit
)

type entry struct {
	title  string
	sub    string
	detail string
}

type model struct {
	screen            screen
	data              UserData
	planTab           int
	cursor            int
	open              int
	width             int
	height            int
	input             string
	inputStep         int
	inputBuf          []string
	now               time.Time
	hasError          bool
	triggerCursor     int
	viewingTriggerId  int
	editingTrigger    bool
	tempTriggerReason string
	commandInput      string
	commandError      string
	relapseNote       string
}

type tickMsg time.Time
