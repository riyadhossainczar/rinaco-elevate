package main

import (
	"fmt"
	"strings"
	"time"
)

func generateFibonacci(n int) []int64 {
	if n <= 0 {
		return []int64{}
	}
	fib := make([]int64, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

var fibSeq = generateFibonacci(30)

func getLevelFromMinutes(minutes int64) int {
	if minutes == 0 {
		return 0
	}
	for i := len(fibSeq) - 1; i >= 0; i-- {
		if minutes >= fibSeq[i] {
			return i
		}
	}
	return 0
}

func getLevelStats(minutes int64) (level int, current, next int64) {
	level = getLevelFromMinutes(minutes)
	current = fibSeq[level]
	if level+1 < len(fibSeq) {
		next = fibSeq[level+1]
	} else {
		next = current + fibSeq[level-1]
	}
	return
}

func currentStart(d UserData) time.Time {
	if len(d.Sessions) == 0 {
		return time.Time{}
	}
	last := d.Sessions[len(d.Sessions)-1]
	if last.EndedAt.IsZero() {
		return last.StartedAt
	}
	return time.Time{}
}






func formatDuration(dur time.Duration) string {
	total := int64(dur.Seconds())
	secs := total % 60
	mins := (total / 60) % 60
	hours := (total / 60 / 60) % 24
	days := (total / 60 / 60 / 24) % 30
	months := (total / 60 / 60 / 24 / 30) % 12
	years := total / 60 / 60 / 24 / 365

	parts := []string{}
	if years > 0 {
		parts = append(parts, fmt.Sprintf("%dy", years))
	}
	if months > 0 {
		parts = append(parts, fmt.Sprintf("%dmo", months))
	}
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if mins > 0 {
		parts = append(parts, fmt.Sprintf("%dm", mins))
	}
	parts = append(parts, fmt.Sprintf("%ds", secs))
	return strings.Join(parts, " ")
}






func getLongestStreak(sessions []Session) time.Duration {
	if len(sessions) == 0 {
		return 0
	}
	var longest time.Duration
	for _, s := range sessions {
		if !s.EndedAt.IsZero() {
			dur := s.EndedAt.Sub(s.StartedAt)
			if dur > longest {
				longest = dur
			}
		}
	}
	return longest
}

func getTotalCleanDays(sessions []Session) int {
	if len(sessions) == 0 {
		return 0
	}
	last := sessions[len(sessions)-1]
	if last.EndedAt.IsZero() {
		return int(time.Since(last.StartedAt).Hours() / 24)
	}
	return 0
}

func getRelapseCount(sessions []Session) int {
	count := 0
	for _, s := range sessions {
		if s.EndedByUs {
			count++
		}
	}
	return count
}

func getDangerousTimeWindow(sessions []Session) (hour int, count int) {
	hourCounts := make(map[int]int)
	for _, s := range sessions {
		if s.EndedByUs {
			h := s.StartedAt.Hour()
			hourCounts[h]++
		}
	}
	maxCount := 0
	for h, c := range hourCounts {
		if c > maxCount {
			maxCount = c
			hour = h
			count = c
		}
	}
	return
}

func getComparisonStats(sessions []Session, now time.Time) (currentMonth, previousMonth int) {
	currentMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, bdTime)
	previousMonthStart := currentMonthStart.AddDate(0, -1, 0)

	for _, s := range sessions {
		if s.EndedByUs {
			if s.StartedAt.After(currentMonthStart) && s.StartedAt.Before(now) {
				currentMonth++
			}
			if s.StartedAt.After(previousMonthStart) && s.StartedAt.Before(currentMonthStart) {
				previousMonth++
			}
		}
	}
	return
}
