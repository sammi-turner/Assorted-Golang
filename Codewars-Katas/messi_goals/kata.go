package main

import "fmt"

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

func main() {
	fmt.Println(Goals(5, 10, 2))
}
