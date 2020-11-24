package main

import (
	"ShirleyGolangTask/models"
	"ShirleyGolangTask/util"
	"testing"
	"time"
)

func TestBuildExecuteTime(t *testing.T) {
      CsvData, _ := util.ReadStats("stats.csv")
      tests := []struct{
      	stats []*models.CsvFields
      	start time.Time
      	end time.Time
      	expectedResult int
	  }{
      	{CsvData, time.Now(), time.Now(), 0},
	  }

	for _, test := range tests{
		actualResult := BuildExecuteTime(test.stats, test.start, test.end)
		if len(actualResult) != test.expectedResult {
			t.Errorf("Shoul return %d items. Returned %d items", test.expectedResult, len(actualResult))
		}
	}
}
