package main

import (
	"ShirleyGolangTask/models"
	"ShirleyGolangTask/util"
	"fmt"
	"log"
	"time"
)

func main() {
	stats, err := util.ReadStats("stats.csv")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Date(2018, 11, 1, 13, 56, 34, 00, time.Local)
	finished := BuildExecuteTime(stats, start, time.Date(2018, 12, 1, 9, 56, 34, 00, time.Local))
	fmt.Printf("%d executed  in a certain time.\n", len(finished))

	users := UserBuildRemoteService(stats)
	fmt.Println("\nTop users:")
	for i, userPair := range users {
		if i == 5 {
			break
		}

		fmt.Printf("%d. User %s: %d builds.\n", i + 1, userPair.Key, userPair.Value)
	}

	//rate := BuildSuccessRate(stats)
	//fmt.Printf("\nSuccess Rate: %.2f%%\n", rate)
}

func BuildExecuteTime(stats []*models.CsvFields, start, end time.Time) []*models.CsvFields {
	predicate := func(stat *models.CsvFields) bool {
		return util.TimeWindow(start, end, stat.TimeRequestFinished)
	}

	return util.FilterStats(stats, predicate)
}


func UserBuildRemoteService(stats []*models.CsvFields) models.PairList {
	ranking := make(map[string]int)

	for _, stat := range stats {
		if _, ok := ranking[stat.UserID]; ok {
			ranking[stat.UserID]++
		} else {
			ranking[stat.UserID] = 1
		}
	}

	return util.RankByValue(ranking)
}


//func BuildSuccessRate(stats []*models.CsvFields) (float64, models.PairList) {
//
//	builds, successful := 0, 0
//
//	for _, stat := range stats {
//		if stat.BuildProcess == 0 {
//			successful++
//
//		builds++
//	}
//
//	successRate := (float64(successful)/float64(builds)) * 100
//	return successRate
//}