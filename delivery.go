package main

type Delivery struct {
	MatchID         int
	Inning          string
	BattingTeam     string
	BowlingTeam     string
	Over            int
	Ball            int
	Batsman         string
	NonStriker      string
	Bowler          string
	IsSuperOver     bool
	WideRuns        int
	ByRuns          int
	LegByRuns       int
	NoBallRuns      int
	PenaltyRuns     int
	BatsmanRuns     int
	ExtraRuns       int
	TotalRuns       int
	PlayerDismissed string
	DismissalKind   string
	Fielder         string
}
