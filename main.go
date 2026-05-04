package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"

	"github.com/golang-sql/civil"
	"github.com/hashicorp/go-set/v3"
)

// Match index positions
const (
	id            = iota // 0
	season               // 1
	city                 // 2
	date                 // 3
	team1                // 4
	team2                // 5
	tossWinner           // 6
	tossDecision         // 7
	result               // 8
	dlApplied            // 9
	winner               // 10
	winsByRun            // 11
	winsByWicket         // 12
	playerOfMatch        // 13
	venue                // 14
	umpire1              // 15
	umpire2              // 16
	umpire3              // 17
)

// Delivery index positions
const (
	matchID         = iota // 0
	inning                 // 1
	battingTeam            // 2
	bowlingTeam            // 3
	over                   // 4
	ball                   // 5
	batsman                // 6
	nonStriker             // 7
	bowler                 // 8
	isSuperOver            // 9
	wideRuns               // 10
	byRuns                 // 11
	legByRuns              // 12
	noBallRuns             // 13
	penaltyRuns            // 14
	batsmanRuns            // 15
	extraRuns              // 16
	totalRuns              // 17
	playerDismissed        // 18
	dismissalKind          // 19
	fielder                // 20
)

const matchesCSVPath = "data/matches.csv"
const deliveriesCSVPath = "data/deliveries.csv"

var csvMatchesFileData []string
var csvDeliveriesFileData []string

func main() {
	//csvFileData := fileReader("data/matches.csv")
	//
	//fmt.Println(csvFileData[1])
	//
	//for _, row := range customCSVSplitter(csvFileData[1]) {
	//	fmt.Println(row)
	//}

	//matches := getMatchSlice()
	//fmt.Println(matches)
	//fmt.Println(matches[1].Venue)

	//deliveries := getDeliverySlice()
	//fmt.Println(deliveries)
	//fmt.Println(deliveries[1].BowlingTeam)

	//fmt.Println(matchIdsOfYear("2015"))
	//fmt.Println(deliveriesOfYear(matchIdsOfYear("2015"))[0])

	matchesPlayedPerYearOfAllTheYears()
	numberOfMatchesWonOfAllTeamsOverAllTheYears()
	extraRunsConcededPerTeamIn2016()
}

func extraRunsConcededPerTeamIn2016() {
	matchIdsOf2016 := matchIdsOfYear("2016")
	deliveriesOf2016 := deliveriesOfYear(matchIdsOf2016)

	runsConcededByEachTeam := make(map[string]int)
	for _, delivery := range deliveriesOf2016 {
		_, exists := runsConcededByEachTeam[delivery.BowlingTeam]
		if exists {
			runsConcededByEachTeam[delivery.BowlingTeam] += delivery.ExtraRuns
		} else {
			runsConcededByEachTeam[delivery.BowlingTeam] = delivery.ExtraRuns
		}
	}

	fmt.Println("Extra runs conceded per team in 2016.")
	for _, key := range slices.Sorted(maps.Keys(runsConcededByEachTeam)) {
		fmt.Printf("%s : %d\n", key, runsConcededByEachTeam[key])
	}
	fmt.Println("---------------------------------------------------")
}

func numberOfMatchesWonOfAllTeamsOverAllTheYears() {
	matches := getMatchSlice()
	matchesMap := make(map[string]int)

	for _, match := range matches {
		_, exists := matchesMap[match.Winner]
		if match.Winner == "" {
			continue
		}
		if exists {
			matchesMap[match.Winner]++
		} else {
			matchesMap[match.Winner] = 1
		}
	}

	fmt.Println("Number of matches won of all teams over all the years of IPL.")
	for team, wins := range matchesMap {
		fmt.Printf("%s : %d\n", team, wins)
	}
	fmt.Println("---------------------------------------------------")
}

func matchesPlayedPerYearOfAllTheYears() {
	matches := getMatchSlice()
	matchesMap := make(map[string]int)

	for _, match := range matches {
		_, exists := matchesMap[match.Season]
		if exists {
			matchesMap[match.Season]++
		} else {
			matchesMap[match.Season] = 1
		}
	}

	fmt.Println("Number of matches played per year of all the years in IPL.")
	for _, key := range slices.Sorted(maps.Keys(matchesMap)) {
		fmt.Printf("%s : %d\n", key, matchesMap[key])
	}
	fmt.Println("---------------------------------------------------")
}

func matchIdsOfYear(season string) set.Set[int] {
	matches := getMatchSlice()
	var filteredmatches set.Set[int]

	for _, match := range matches {
		if match.Season == season {
			filteredmatches.Insert(match.ID)
		}
	}

	return filteredmatches
}

func deliveriesOfYear(matchIDs set.Set[int]) []Delivery {
	deliveries := getDeliverySlice()
	var filtereddeliveries []Delivery

	for _, delivery := range deliveries {
		if matchIDs.Contains(delivery.MatchID) {
			filtereddeliveries = append(filtereddeliveries, delivery)
		}
	}

	return filtereddeliveries
}

func getMatchSlice() []Match {
	var matches []Match

	if csvMatchesFileData == nil {
		csvMatchesFileData = fileReader(matchesCSVPath)
	}

	for _, row := range csvMatchesFileData {
		splitRowData := customCSVSplitter(row)

		id, _ := strconv.Atoi(splitRowData[id])
		date, _ := civil.ParseDate(splitRowData[date])
		winsByRun, _ := strconv.Atoi(splitRowData[winsByRun])
		winsByWicket, _ := strconv.Atoi(splitRowData[winsByWicket])

		match := Match{
			ID:            id,
			Season:        splitRowData[season],
			City:          splitRowData[city],
			Date:          date,
			Team1:         splitRowData[team1],
			Team2:         splitRowData[team2],
			TossWinner:    splitRowData[tossWinner],
			TossDecision:  splitRowData[tossDecision],
			Result:        splitRowData[result],
			DLApplied:     splitRowData[dlApplied] == "1",
			Winner:        splitRowData[winner],
			WinsByRun:     winsByRun,
			WinsByWicket:  winsByWicket,
			PlayerOfMatch: splitRowData[playerOfMatch],
			Venue:         splitRowData[venue],
			Umpire1:       splitRowData[umpire1],
			Umpire2:       splitRowData[umpire2],
			Umpire3:       splitRowData[umpire3],
		}
		matches = append(matches, match)
	}

	return matches
}

func getDeliverySlice() []Delivery {
	var deliveries []Delivery

	if csvDeliveriesFileData == nil {
		csvDeliveriesFileData = fileReader(deliveriesCSVPath)
	}

	for _, row := range csvDeliveriesFileData {
		splitRowData := customCSVSplitter(row)

		matchID, _ := strconv.Atoi(splitRowData[matchID])
		over, _ := strconv.Atoi(splitRowData[over])
		ball, _ := strconv.Atoi(splitRowData[ball])
		wideRuns, _ := strconv.Atoi(splitRowData[wideRuns])
		byRuns, _ := strconv.Atoi(splitRowData[byRuns])
		legByRuns, _ := strconv.Atoi(splitRowData[legByRuns])
		noBallRuns, _ := strconv.Atoi(splitRowData[noBallRuns])
		penaltyRuns, _ := strconv.Atoi(splitRowData[penaltyRuns])
		batsmanRuns, _ := strconv.Atoi(splitRowData[batsmanRuns])
		extraRuns, _ := strconv.Atoi(splitRowData[extraRuns])
		totalRuns, _ := strconv.Atoi(splitRowData[totalRuns])

		delivery := Delivery{
			MatchID:         matchID,
			Inning:          splitRowData[inning],
			BattingTeam:     splitRowData[battingTeam],
			BowlingTeam:     splitRowData[bowlingTeam],
			Over:            over,
			Ball:            ball,
			Batsman:         splitRowData[batsman],
			NonStriker:      splitRowData[nonStriker],
			Bowler:          splitRowData[bowler],
			IsSuperOver:     splitRowData[isSuperOver] == "1",
			WideRuns:        wideRuns,
			ByRuns:          byRuns,
			LegByRuns:       legByRuns,
			NoBallRuns:      noBallRuns,
			PenaltyRuns:     penaltyRuns,
			BatsmanRuns:     batsmanRuns,
			ExtraRuns:       extraRuns,
			TotalRuns:       totalRuns,
			PlayerDismissed: splitRowData[playerDismissed],
			DismissalKind:   splitRowData[dismissalKind],
			Fielder:         splitRowData[fielder],
		}
		deliveries = append(deliveries, delivery)
	}

	return deliveries
}

func customCSVSplitter(dataRow string) []string {
	var result []string
	var field []rune
	quotesActive := false

	for _, c := range dataRow {
		if c == '"' {
			quotesActive = !quotesActive
		} else if c == ',' && !quotesActive {
			result = append(result, string(field))
			field = []rune{}
		} else {
			field = append(field, c)
		}
	}

	result = append(result, string(field))

	return result
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
	return csvFileData[1:]
}
