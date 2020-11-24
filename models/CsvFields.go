package models

import "time"

type CsvFields struct {
	BuildID             string
	UserID              string
	TimeRequestReceived time.Time
	TimeRequestBegan    time.Time
	TimeRequestFinished time.Time
	BuildDeleted        bool
	BuildProcess        int
	ImageSize           int
}