package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Match index positions
	const (
		id            = 0
		season        = 1
		city          = 2
		date          = 3
		team1         = 4
		team2         = 5
		tossWinner    = 6
		tossDecision  = 7
		result        = 8
		dlApplied     = 9
		winner        = 10
		winsByRun     = 11
		winsByWicket  = 12
		playerOfMatch = 13
		venue         = 14
		umpire1       = 15
		umpire2       = 16
		umpire3       = 17
	)

	// Delivery index positions
	const (
		matchID         = 0
		inning          = 1
		battingTeam     = 2
		bowlingTeam     = 3
		over            = 4
		ball            = 5
		batsman         = 6
		nonStriker      = 7
		bowler          = 8
		isSuperOver     = 9
		wideRuns        = 10
		byRuns          = 11
		legByRuns       = 12
		noBallRuns      = 13
		penaltyRuns     = 14
		batsmanRuns     = 15
		extraRuns       = 16
		totalRuns       = 17
		playerDismissed = 18
		dismissalKind   = 19
		fielder         = 20
	)

	csvFileData := fileReader("data/matches.csv")

	fmt.Println(csvFileData[1])
}

func fileReader(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
	}(file)

	var csvFileData []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		csvFileData = append(csvFileData, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	return csvFileData
}
