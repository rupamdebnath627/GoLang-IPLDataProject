package main

import "time"

type Match struct {
	ID            int
	Season        string
	City          string
	Date          time.Time
	Team1         string
	Team2         string
	TossWinner    string
	TossDecision  string
	Result        string
	DLApplied     bool
	Winner        string
	WinsByRun     int
	WinsByWicket  int
	PlayerOfMatch string
	Venue         string
	Umpire1       string
	Umpire2       string
	Umpire3       string
}
