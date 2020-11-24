package util

import (
	"ShirleyGolangTask/models"
	"encoding/csv"
	"os"
	"sort"
	"strconv"
	"time"
)

func TimeWindow(start, end, inspect time.Time)bool  {
	return inspect.After(start) && inspect.Before(end)
}


// FilterStats returns a slice of all statistics that satisfy predicate f
func FilterStats(stats []*models.CsvFields, f func(statistic *models.CsvFields) bool) []*models.CsvFields {
	filtered := make([]*models.CsvFields, 0)
	for _, stat := range stats {
		if f(stat) {
			filtered = append(stats, stat)
		}
	}

	return filtered
}

func ReadStats(filename string) (stats []*models.CsvFields, err error) {
	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		received, _ := time.Parse(time.RFC3339, line[2])
		began, _ := time.Parse(time.RFC3339, line[3])
		finished, _ := time.Parse(time.RFC3339, line[4])
		deleted, _ := strconv.ParseBool(line[5])
		process, _ := strconv.Atoi(line[6])
		size, _ := strconv.Atoi(line[7])

		stat := &models.CsvFields{
			BuildID: line[0],
			UserID : line[1],
			TimeRequestReceived: received,
			TimeRequestBegan: began,
			TimeRequestFinished: finished,
			BuildDeleted: deleted,
			BuildProcess: process,
			ImageSize: size,
		}

		stats = append(stats, stat)
	}

	return
}

func RankIntByMapValue(values map[int]int) models.PairList  {
	pl := make(models.PairList, len(values))
	i := 0

	for k, v := range values {
		pl[i] = models.Pair{Key: k, Value: v}
		i++
	}
	sort.Sort(pl)
	return pl
}

func RankByValue(values map[string]int) models.PairList {
	pl := make(models.PairList, len(values))
	i := 0

	for k, v := range values {
		pl[i] = models.Pair{Key: k, Value: v}
		i++
	}

	sort.Sort(pl)
	return pl
}