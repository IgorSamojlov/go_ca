package catalog

import "time"

type FundID int64

type Fund struct {
	ID                FundID
	Code              string
	Title             string
	CodeOrder         string
	NumberOfArchFiles int
	StartYear         int
	EndYear           int
	ArchiveID         int64
	UserID            int64
	FundType          int
	DSP               bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
